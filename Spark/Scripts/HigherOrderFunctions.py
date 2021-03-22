### HIGHER ORDER FUNCTIONS ###
# arrays and maps/dicts allow a range of higher order functions
# e.g. union, select_distinct, explode etc.
# you can find details of these on p.139 of the LearningSpark2.0 book
# here is a good website for understanding functions like explode
# https://sparkbyexamples.com/spark/explode-spark-array-and-map-dataframe-column/

# main method
if __name__ == "__main__":

    # load libs
    from pyspark.sql import SparkSession
    from pyspark.sql.types import *
    from pyspark.sql.functions import expr
    
    # create session
    spark = (SparkSession
             .builder
             .appName('functions')
             .getOrCreate())
    
    # create a simple dataframe with 2 rows, each containing an array
    schema = StructType(StructField("celsius", ArrayType(IntegerType())))
    data = [[35, 36, 32, 37, 30, 31], [28, 30, 26, 36, 33]]
    df = spark.createDataFrame(data, schema)
    df.createOrReplaceTempView('table')
    
    ### LAMBDA EXPRESSIONS ###
    # these allow you to pass an array of input values to an anonymous function
    # which applies the function to each value and returns the transformed results
    # TRANSFORM lets you apply a formula/equation to your values (like lambda)
    spark.sql("""SELECT celsius,
                 transform(celsius, t -> ((t * 9) div 5) + 32) as fahrenheit
                 FROM table""").show()
                 
    # FILTER returns an array of values where the values in your array match
    # the specified criteria (i.e. here a list of temps above 38)
    spark.sql("""SELECT celsius,
                 filter(celsius, t -> t > 38) AS high
                 FROM table""").show()
    
    # EXISTS returns boolean true if the specified value(s) exists in your array
    spark.sql("""SELECT celsius,
                 exists(celsius, t -> t = 38)
                 FROM table""")
                
    # other built in functions can be found here
    # https://spark.apache.org/docs/latest/api/sql/index.html
                 
    # use CAST to convert column types
    df = (df
          .withColumn('delay', expr("CAST(delay as INT) as delay"))
          .withColumn('distance', expr('CAST(delay as INT) as distance')))
    
    # create temp view of above to enable SQL
    df.createOrReplaceTempView('table')
    
    # filter with expr and date like to convert datetime format
    view = (df
            .filter(expr("""expr(origin == 'SEA', and destination == 'SFO'
                      and date like '01010%' and delay > 0""")))
     
    # create temp view of above filtered df to enable SQL
    # this new table is now much smaller than the original, making it faster
    view.getOrCreateTempView('view')
    
    # extract data from much smaller, filtered table (faster, efficient)
    spark.sql("""SELECT * FROM view LIMIT 10""").show()
    
    ### UNION ###
    # this is used to join two dataframes with the same schema
    # here we create a new_df as the union of df and view
    # this will lead to the duplication of data
    new_df = df.union(view)
    
    ### JOIN ###
    # joins connect two dataframes based on certain columns
    # the default is an inner join (i.e. shared values) but other joins include
    # inner, cross, outer, full, full_outer, left, left_outer, right,
    # right_outer, left_semi, and left_anti
    spark.sql("""SELECT a.name, a.age, b.address, b.income
                 FROM df a
                 JOIN df_2 b
                 ON a.id = b.id""").show()
                 
    ### WINDOWING ###
    # a window is a subset of rows from a table where you can or have performed
    # custom actions such as ranking, ntile splitting, % rank etc. on the rows
    # an example of a simple ranking window action is below, but further detail
    # on windows can be found here
    # https://databricks.com/blog/2015/07/15/introducing-window-functions-in-spark-sql.html
    # select a subset of data from our main data (3 origins, 7 destinations)
    spark.sql("""CREATE TABLE sub_delays AS
                 SELECT origin, destination, sum(delay) AS TotalDelays
                 FROM flight_delays
                 WHERE origin IN ('SEA', 'SFO', 'JFK')
                 AND destination IN ('SEA', 'SFO', 'JFK', 'DEN',
                                     'ORD', 'LAX', 'ATL')
                 GROUP BY origin, destination""")
                 
    # if you wanted to see the top 3 delayed routes for each origin airport
    # you could write the following query 3 times (1 for each origin)
    spark.sql("""SELECT origin, destination, TotalDelays
                 FROM sub_delays
                 WHERE origin = 'origin_name'
                 GROUP BY origin, destination
                 ORDER BY TotalDelays DESC
                 LIMIT 3""")
                 
    # but this is inefficient, therefore you could use the window function
    # below to do this in a more elegant manner (implementing dense_rank())
    # this is basically a way of filtering, sorting and iterating all at once
    # NOTE: each window operation (i.e. each origin rank chunk) will get sent
    # to a single partition when executed, so it's important to ensure your
    # tables aren't unbounded (i.e. specific, small range of rows only)
    spark.sql("""SELECT origin, destination, TotalDelays, rank
                 FROM (SELECT origin, destination, TotalDelays, dense_rank()
                       OVER (PARTITION BY origin ORDER BY TotalDelays DESC)
                       AS rank
                       FROM sub_delays) t
                 WHERE rank <= 3)""").show()
    
    ### MODIFYING DATAFRAMES ###
    # dataframes are immutable, but you can create new versions of dataframes
    # which you can modify (in the Spark lineage) via certain modifications
    # ADDING COLUMNS
    foo2 = foo2.withColumn('status',
                           expr('CASE WHEN delay <=10 THEN "on-time" \
                                 ELSE "delayed" END'))
                                 
    # DROPPING COLUMNS
    foo3 = foo2.drop('delay')

    # RENAMING COLUMNS
    foo4 = foo3.withColumnRenamed('status', 'flight_check')
    
    # PIVOTING TABLES (i.e. switch cols with rows)
    # here we take a single month column (with number for months i.e. 1, 2)
    # and pivot so that the columns become months (i.e. Jan, Feb across)
    # and we also create 2 columns per month, 1 for avg and 1 for max delay
    spark.sql("""SELECT * FROM (
                SELECT destination, month, delay
                FROM flight_table WHERE origin = 'SFO'
                )
                PIVOT (
                CAST(avg(delay)) AS DECIMAL(4, 2) AS avg_delay,
                CAST(max(delay) AS max_delay)
                FOR month IN (1 JAN, 2 FEB))""")
                
    # NOTE: it's always best to try and use DSL with Spark so that you're
    # implementing the DataFrame (or Dataset) specific code as opposed to
    # writing opaque functions (e.g. lambda, UDFs) so that Spark can see into
    # your functions and optimize them rather than converting to DSL at runtime
    
    # stop session
    spark.stop()
    