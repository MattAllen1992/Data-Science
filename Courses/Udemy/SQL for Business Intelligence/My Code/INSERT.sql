# INSERT INTO VALUES used to add new rows
# specify columns to insert data into and relevant values
# value order must match column order specified
# don't have to specify all columns in the table (providing defaults etc. handle the rest)
# check column data types (in table/column info) to ensure you're entering the correct data types
# SQL will convert strings into ints etc. if necessary but this takes time and is not best practice
insert into employees (emp_no, birth_date, first_name, last_name, gender, hire_date)
values (999901, '1992-07-03', 'Matthew', 'Allen', 'M', '2018-01-01');

# you don't have to specify column names and orders
# but the values order must then match the default order of columns in the table
insert into employees
values (999901, '1992-07-03', 'Matthew', 'Allen', 'M', '2018-01-01');

# insert information from existing table into another
insert into departments_dup
select *
from departments;