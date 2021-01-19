#!/usr/bin/env python
# coding: utf-8

# In[1]:


# load libraries
import numpy as np
import pandas as pd
import pickle
from sklearn.preprocessing import StandardScaler
from sklearn.base import BaseEstimator, TransformerMixin

# create custom scaler class
class CustomScaler(BaseEstimator, TransformerMixin):
    
    # constructor to instantiate class objects
    def __init__(self, columns, copy=True, with_mean=True, with_std=True):
        self.scaler = StandardScaler(copy, with_mean, with_std)
        self.columns = columns
        self.mean_ = None
        self.var_ = None
    
    # fit model to input data, calculate mean and var
    def fit(self, X, y=None):
        self.scaler.fit(X[self.columns], y)
        self.mean_ = np.mean(X[self.columns])
        self.var_ = np.var(X[self.columns])
        return self
    
    # transform raw data to scaled/standardized
    def transform(self, X, y=None, copy=None):
        init_col_order = X.columns
        X_scaled = pd.DataFrame(self.scaler.transform(X[self.columns]), columns=self.columns)
        X_not_scaled = X.loc[:, ~X.columns.isin(self.columns)]
        return pd.concat([X_not_scaled, X_scaled], axis=1)[init_col_order]
    
# create special class to predict new data
class absenteeism_model():
    
    # constructor to instantiate class objects
    def __init__(self, model_file, scaler_file):
        # read in model and scaler files
        with open('model', 'rb') as model_file, open('scaler', 'rb') as scaler_file:
            self.reg = pickle.load(model_file)
            self.scaler = pickle.load(scaler_file)
            self.data = None
            
    # load data as csv and pre-process
    def load_and_clean_data(self, data_file):
        # import data
        df = pd.read_csv(data_file, delimiter=',')
        # store data in var
        self.df_with_predictions = df.copy()
        # drop ID col
        df = df.drop(['ID'], axis = 1)
        # preserve structure of df
        df['Absenteeism Time in Hours'] = 'NaN'
        
        # new df with dummy vars for sickness reasons
        reason_columns = pd.get_dummies(df['Reason for Absence'], drop_first = True)
        
        # split reasons into 4 groups
        reason_type_1 = reason_columns.loc[:, 1:14].max(axis=1)
        reason_type_2 = reason_columns.loc[:, 15:17].max(axis=1)
        reason_type_3 = reason_columns.loc[:, 18:21].max(axis=1)
        reason_type_4 = reason_columns.loc[:, 22:].max(axis=1)
        
        # avoid multicollinearity, drop raw absence reasons
        df = df.drop(['Reason for Absence'], axis=1)
        
        # concatenate dummy vars
        df = pd.concat([df, reason_type_1, reason_type_2, reason_type_3, reason_type_4], axis=1)
        
        # rename vars
        column_names = ['Date', 'Transportation Expense', 'Distance to Work', 'Age',
                        'Daily Work Load Average', 'Body Mass Index', 'Education', 'Children',
                        'Pets', 'Absenteeism Time in Hours', 'Reason_1', 'Reason_2', 'Reason_3', 'Reason_4']
        df.columns = column_names
        
        # re-order cols
        column_names_reordered = ['Reason_1', 'Reason_2', 'Reason_3', 'Reason_4', 'Date', 'Transportation Expense',
                                  'Distance to Work', 'Age', 'Daily Work Load Average', 'Body Mass Index', 'Education',
                                  'Children', 'Pets', 'Absenteeism Time in Hours']
        df = df[column_names_reordered]
        
        # convert date to datetime
        df['Date'] = pd.to_datetime(df['Date'], format='%d%m%Y')
        
        # create list of months
        list_months = []
        for i in range(df.shape[0]):
            list_months.append(df['Date'][i].month)
            
        # store in df
        df['Month Value'] = list_months
        
        # create day of the week
        df['Day of the Week'] = df['Date'].apply(lambda x: x.weekday())
        
        # drop date from df
        df = df.drop(['Date'], axis=1)
        
        # re-order cols
        column_names_upd = ['Reason_1', 'Reason_2', 'Reason_3', 'Reason_4', 'Month Value', 'Day of the Week',
                            'Transportation Expense', 'Distance to Work', 'Age', 'Daily Work Load Average',
                            'Body Mass Index', 'Education', 'Children', 'Pets', 'Absenteeism Time in Hours']
        df = df[column_names_upd]
        
        # map education as dummy var
        df['Education'] = df['Education'].map({1:0, 2:1, 3:1, 4:1})
        
        # replace NaN
        df = df.fillna(value=0)
        
        # drop absenteeism
        df = df.drop(['Absenteeism Time in Hours'], axis=1)
        
        # drop redundant vars (backward elimination)
        df = df.drop(['Day of the Week', 'Daily Work Load Average', 'Distance to Work'], axis=1)
        
        # checkpoint data as fully pre-processed df
        self.preprocessed_data = df.copy()
        
        # store in class var for later use
        # transform using fitted scaler
        self.data = self.scaler.transform(df)
        
    # function to output predicted probabilities
    def predicted_probability(self):
        if (self.data is not None):
            pred = self.reg.predict_proba(self.data)[:,1] # output class 1 only
            return pred
        
    # function to output predicted result (i.e. 0 or 1 instead of probs)
    def predicted_output_category(self):
        if (self.data is not None):
            pred_outputs = self.reg.predict(self.data)
            return pred_outputs
        
    # predict outputs and probabilities, append to new df
    def predicted_outputs(self):
        if (self.data is not None):
            self.preprocessed_data['Probability'] = self.reg.predict_proba(self.data)[:,1]
            self.preprocessed_data['Prediction'] = self.reg.predict(self.data)
            return self.preprocessed_data

