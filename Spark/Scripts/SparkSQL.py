# session libs and query functions
from pyspark.sql import SparkSession
from pyspark.sql.functions import col, desc, to_timestamp, month, when

# main method
if __name__ == "__main__":
    
    # create session
    spark = (SparkSession
             .builder
             .appName('SparkSQL')
             .getOrCreate())
    
    # define schema
    schema = ("date STRING, delay INT, \
              distance INT, origin STRING, destination STRING")

    # load data into df
    df = spark.read.csv('path/to/file', schema=schema, header=True)
    
    # create temp view
    # enables direct use of SQL code on this object
    # we can refer to 'tempView' as a table in later steps
    df.createOrReplaceTempView('tempView')
    
    # select all flights whose distance is > 1000 miles
    spark.sql("""SELECT origin, destination, distance
                 FROM tempView
                 WHERE distance > 1000
                 ORDER BY distance DESC""").show(10)
                 
    # identical result to above except using DataFrame API code
    (df.select('origin', 'destination', 'distance')
       .where(col('distance') > 1000)
       .orderBy(col('distance'), ascending=False)).show(10)
    
    # identical result to above except using alternate code
    (df.select('origin', 'destination', 'distance')
       .where('distance > 1000')
       .orderBy(desc('distance'))).show(10)
                 
    # select flights between SFO and ORD with >2h delay
    spark.sql("""SELECT date, origin, destination, delay
                 FROM tempView
                 WHERE delay > 120
                 AND origin = 'SFO' AND destination = 'ORD'
                 ORDER BY delay DESC""").show(10)
                 
    # same as above but using DataFrame API
    (df.select('date', 'origin', 'destination', 'delay')
       .where((col('delay') > 120) &
              (col('origin') == 'SFO') &
              (col('destination') == 'ORD'))
       .orderBy(desc('delay'))).show(10)
                 
    # convert date to actual datetime (101 = "mm/dd/yyyy")
    # show total delays between SFO and ORD by month
    spark.sql("""SELECT MONTH(CONVERT(datetime, date, 101)) AS month,
                        origin, destination, SUM(delay)
                 FROM tempView
                 WHERE delay > 120
                 AND origin = 'SFO' AND destination = 'ORD'
                 GROUPBY month""").show(10)
                 
    # same as above but using DataFrame API
    (df.select(to_timestamp(col('date'), 'yyyy-MM-dd HH:mm:ss').alias('datetime'),
               'origin', 'destination', sum('delay'))
       .where((col('delay') > 120) &
              (col('origin') == 'SFO') &
              (col('destination') == 'ORD'))
       .groupBy(month('datetime').alias('month')))
                 
    # use CASE for conditional operations
    # create a new column flight_delay to track delay category
    spark.sql("""SELECT delay, origin, destination
                 CASE
                     WHEN delay >= 360 THEN 'EXTREME DELAY'
                     WHEN delay >= 120 AND delay < 360 THEN 'LONG DELAY'
                     WHEN delay >= 60 AND delay < 120 THEN 'DELAY'
                     WHEN delay >= 0 AND delay < 60 THEN 'MINOR DELAY'
                     WHEN delay = 0 THEN 'NO DELAY'
                     ELSE 'EARLY'
                 END AS flight_delay
                 FROM tempView
                 ORDER BY origin, delay DESC""").show(10)
                 
    # same as above but using DataFrame API
    (df.select('delay', 'origin', 'destination')
       .when(col('delay') >= 360, 'EXTREME DELAY')
       .when((col('delay') >= 120) & (col('delay') < 360), 'LONG DELAY')
       .when((col('delay') >= 60) & (col('delay') < 120), 'DELAY')
       .when((col('delay') >= 0) & (col('delay') < 60), 'MINOR DELAY')
       .when(col('delay') == 0, 'ON TIME')
       .otherwise('EARLY').alias('flight_delay')
       .orderBy(['origin', 'delay'], ascending=False)).show(10)