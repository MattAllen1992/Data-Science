# load libs
from pyspark.sql import SparkSession
from pyspark.sql.types import LongType
from pyspark.sql.functions import col, pandas_udf

# main method
if __name__ == "__main__":
    
    # create session
    spark = (SparkSession
             .builder
             .appName('UDF')
             .getOrCreate())
    
    ### UDFS ###
    # UDFs allow you to write custom functions in whatever language you like
    # (even if the specific functions etc. aren't compatible) and then use
    # them within Spark as if you were writing them in native Python for example
    # create function which cubes value
    def cubed(x):
        return x * x * x
    
    # register UDF with Spark (3rd part is return type)
    spark.udf.register('cubed', cubed, LongType())
    
    # create temp SQL view
    # range(1,9) creates a df with a col called 'id' with vals from 1 to 9
    spark.range(1, 9).createOrReplaceTempView('temp_view')
    
    # use cubed function in Spark SQL query
    spark.sql("""SELECT id,
                        cubed(id) AS id_cubed
                 FROM temp_view""").show()
                 
    # NOTE: UDFs are a black box to Spark, meaning it can't optimize them, so
    # only use them if you absolutely have to, otherwise you'll lose performance
                 
    # one way to get around this is to use Spark's 'pandas_utf' which ensures
    # that you're working with vectors (i.e. full cols) to improve performance
    # create a cubed function with explicit in/out types
    # this function takes a Pandas series
    def cubed(a: pd.Series) -> pd.Series:
        return a * a * a
    
    # create a pandas UDF for the cubed function
    cubed_udf = pandas_udf(cubed, returnType=LongType())
    
    # this line will then run the optimized process where Spark can optimize
    # this will create spark jobs (i.e. granular) which are more efficient
    df.select("id", cubed_udf(col("id"))).show()
    
    # as opposed to this local call which would be much slower/worse
    # this will not create jobs and will only interact with the driver
    print(cubed(x))