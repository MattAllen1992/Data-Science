# Indexes can be used to speed up searching for items in a table (when querying)
# Essentially you create an index of the columns you're searching in and this enables faster queries
# Index objects are created when you do this and these take up space in your DB, therefore if you're working
# with a small DB then the cost of creating an index object might outweight the speed improvements of an index search,
# whereas with big DBs the index object will speed up the search dramatically.
# You can create individual index objects which simply index a single column, or you can create a composite index
# which indexes more than one column (examples of each below).

# create an indexed hire_date column
CREATE INDEX i_hire_date ON employees(hire_date);

# the below select query will now run far quicker due to the indexed column involved in the query
SELECT * FROM employees WHERE hire_date > '2000-01-01';

# create a composite index on more than one column
CREATE INDEX i_composite ON employees(first_name, last_name);

# the below query will now run far quicker due to the composite index created above
SELECT * FROM employees WHERE first_name = 'Georgi' AND last_name = 'Facello';

# this will show all created index objects from this table from this database
# you can also do this by investigating the db or table objects and going to their "indexes" tab
SHOW INDEX FROM employees FROM employees;