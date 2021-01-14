# create DB from scratch (replace if already exists)
DROP DATABASE IF EXISTS predicted_outputs;
CREATE DATABASE IF NOT EXISTS predicted_outputs;

# select DB for use
USE predicted_outputs;

# create table matching data input from Python
# BIT stands for binary data
DROP TABLE IF EXISTS predicted_outputs;
CREATE TABLE predicted_outputs
(
	Reason_1 BIT NOT NULL,
    Reason_2 BIT NOT NULL,
    Reason_3 BIT NOT NULL,
    Reason_4 BIT NOT NULL,
    month_value INT NOT NULL,
    transportation_expense INT NOT NULL,
    age INT NOT NULL,
    body_mass_index INT NOT NULL,
    education BIT NOT NULL,
    children INT NOT NULL,
    pet INT NOT NULL,
    probability FLOAT NOT NULL,
    prediction BIT NOT NULL
);

# select everything from table
SELECT
	*
FROM
	predicted_outputs;