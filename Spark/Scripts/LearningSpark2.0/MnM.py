# find spark
import findspark
findspark.init('/home/matt/Spark')

# load libs
import sys
import traceback
from pyspark.sql import SparkSession
from pyspark.sql.functions import count

# main method
if __name__ == "__main__":
    # create dummy input argv
    in_arg = '/home/matt/Spark/data/mnm_dataset.csv'
        
    # create spark session
    spark = (SparkSession
             .builder
             .appName('PythonMnMCount')
             .getOrCreate())
    
    # get data from command line args
    mnm_file = in_arg[1]
    
    # load csv data into Spark DataFrame
    mnm_df = (spark.read.format('csv')
              .option('inferSchema', 'true')
              .option('header', 'true')
              .load(mnm_file))
    
    # group by state and color, show count per row
    count_mnm_df = (mnm_df
                    .select('State', 'Color', 'Count')
                    .groupBy('State', 'Color')
                    .agg(count('Count').alias('Total'))
                    .orderBy('Total', ascending=False))
    
    # show top 60 rows
    # note: show() is an action, therefore above transformations
    # will be triggered at this point
    count_mnm_df.show(n=60, truncate=False)
    print("Total rows: %d", (count_mnm_df.count()))
    
    # stop session once finished
    spark.stop()