### SETUP SPARK ###
# load libs
from pyspark.sql import SparkSession

# main method
if __name__ == "__main__":

    # create session
    spark = (SparkSession
             .builder
             .appName('DatabasesTablesViews')
             .getOrCreate())
    
    ### CREATE CUSTOM DATABASE ###
    # create new database
    # by default Spark creates and uses 'default' database
    # the below code lets you create your own custom db
    spark.sql('CREATE DATABASE my_db')
    
    # tell spark to use this db
    # any code from now will automatically use this db
    spark.sql('USE my_db')
    
    ### MANAGED TABLES ###
    # managed tables are managed by Spark
    # meaning that Spark stores both the table and the metadata
    # metadata contains things like physical location, description, db name etc.
    # define schema
    schema = ("date STRING, delay INT, \
               distance INT, origin STRING, destination STRING")
    
    # load data into df
    df = spark.read.csv('path/to/file', schema=schema, header=True)
    
    # create managed table using above file and schema
    df.write.saveAsTable('managed_table')
    
    # same as above 3 lines but using SQL format
    spark.sql("""CREATE TABLE managed_table (date STRING delay INT \
                 origin STRING destination STRING)""")
    
    ### UNMANAGED TABLES ###
    # unmanaged tables only store their metadata in Spark
    # the tables themseleves are handled by an external source (e.g. Cassandra)
    (df.write
       .option('path', 'path/to/file')
       .saveAsTable('unmanaged_table'))
       
    # same as above but using SQL format
    spark.sql("""CREATE TABLE unmanaged_data (date STRING delay INT \
                 origin STRING destination STRING)
                 USING csv OPTIONS  (PATH 'path/to/file')""")
    
    ### VIEWS ###
    # views are temporary versions of tables
    # views do not store actual data in the DB
    # once the session closes they are gone
    # create query to create view
    df_sfo = spark.sql("""SELECT date, delay, origin, destination
                           FROM table WHERE origin='SFO'""")
                           
    # create view using above query and name it
    # global views can be accessed between multiple Spark sessions
    # standard views are visible within the current session only
    df_sfo.createOrReplaceGlobalTempView('df_temp_view')
    df_sfo.createOrReplaceTempView('df_temp_view')
    
    # you can then read these views as tables
    spark.read.table('df_temp_view')
    
    # and pass SQL queries to them
    spark.sql("""SELECT * FROM df_temp_view""")
    
    ### METADATA CATALOG ###
    # Spark stores metadata for your databases, tables etc.
    # you can access it using the 'catalog' element of your session
    # a few examples below
    spark.catalog.listDatabases()
    spark.catalog.listTables()
    spark.catalog.listColumns('df_temp_view')
    
    ### CACHING TABLES ###
    # Spark caches tables and other data elements
    # it does this to prevent having to read the original file each time
    # which can be expensive in terms of time, memory and processing
    # Spark allows you to lazily cache tables meaning they are only
    # cached when they are first used instead of immediately, decreasing initial load