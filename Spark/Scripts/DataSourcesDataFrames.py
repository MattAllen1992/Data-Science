### DATAFRAME READER ###
# this allows us to load a variety of data sources into a Spark DataFrame
# the below line is the standard format for all the reader options
# format defines the data format (e.g. csv, json, text)
# options are custom for different data types and require key/value pairs
# schema allows you to infer schema or specify manually
# load requires the path to the data source itself
DataFrameReader.format(args).option("key", "value").schema(args).load(path)

# Parquet format is the preferred option for Spark because it is compressed,
# optimized and comes with pre-defined metadata (allowing inference)

### DATAFRAME WRITER ###
# this writes a dataframe's contents to a specific data format
# below is the standard syntax for its usage
# format and option are the same as above and save is simply a path
# bucket/partition by determine which column chunks the df will be split on
DataFrameWriter.format(args)
               .option(args)
               .bucketBy(args)
               .partitionBy(args)
               .save(path)

### PARQUET TABLES ###
# Parquet tables are the default for Spark because they are compressed
# meaning rows and cols are reduced and querying is far quicker
# they are stored in folders with the data and metadata which means you don't
# need to specify a schema for static data (although streaming requires one)
# this code shows how to read a parquet table into a dataframe
df = spark.read.format('parquet').load('path/to/parquet/folder')

# once you've saved the df into a table, you can write SQL queries for it
df.write.saveAsTable('table')

# sql queries for imported parquet table
spark.sql("""SELECT * FROM table""").show()

# writing a parquet table creates a folder with the data and metadata files
# because parquet is the defauly Spark format, you could omit 'parquet' below
(df.write.format('parquet')
         .mode('overwrite')
         .option('compression', 'snappy')
         .save('path/to/folder'))

# you can also just write your table directly to a SQL table
(df.write
   .mode('overwrite')
   .saveAsTable('table'))

### JSON (JAVASCRIPT OBJECT NOTATION) TABLES ###
# json is essentially an easier to read and parse version of XML
# it can be either single or multi-line in format and Spark lets you handle
# either using the 'multiline=True' parameter when reading in
# here's how to read in json
df = spark.read.format("json").load(path)

# and here's how to write out json
(df.write.format("json")
   .mode("overwrite")
   .option("compression", "snappy")
   .save("/tmp/data/json/df_json"))

### CSV ###
# CSV (comma-separated value) files are a popular format and easily read in
# there are many options for different separators, escape chars etc. in the docs
schema = "DEST_COUNTRY_NAME STRING, ORIGIN_COUNTRY_NAME STRING, count INT"
df = (spark.read.format("csv")
                .option("header", "true")
                .schema(schema)
                .option("mode", "FAILFAST") # exit if error(s)
                .option("nullValue", "") # replace nulls with quotes
                .load(path))

# writing out is also straightforward
df.write.format("csv").mode("overwrite").save("path/to/file")

### OTHER FORMATS ###
# There are many other input formats that Spark can handle (listed below)
# for full details on how to use them, refer to the Spark docs
# 1) Avro (Apache Kafka) is for speedily serializing/deserializing message data
# 2) ORC is an optimized format for when you really need to reduce CPU usage for
# intensive operations (e.g. filters, scans, aggregations etc.)
# 3) Images can be handled for ML and deep learning (e.g. computer vision)
# 4) Binary data can also be handled