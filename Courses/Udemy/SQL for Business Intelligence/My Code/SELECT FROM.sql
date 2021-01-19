# select specific columns
SELECT first_name, last_name FROM employees.employees;

# select all columns
SELECT * FROM employees.empoyees;

# select with specific conditions
SELECT * FROM employees WHERE first_name = 'Denis';

# select with multiple conditions
SELECT * FROM employees.employees WHERE (first_name = 'Kellie' OR first_name = "Elvis") AND gender = 'M';

# NOTE: AND is prioritised/read first over OR
# below code gives a different result to the above code
SELECT * FROM employees.employees WHERE first_name = 'Kellie' OR first_name = "Elvis" AND gender = 'M';

# select when condition matches list contents
# NOTE: this is quicker than chained OR statements (check time to run in 'action output' below)
SELECT * FROM employees.employees WHERE first_name IN ('Cathie', 'Mark', 'Nathan');
SELECT * FROM employees.employees WHERE first_name NOT IN ('Cathie', 'Mark', 'Nathan');

# pattern matching (~regex)
# wildcards: %, _ and * (i.e. means you can put anything in this place)
# NOTE: this is case insensitive
SELECT * FROM employees.employees WHERE first_name LIKE ('_ark'); # 4 letters only, ending in 'ark'
SELECT * FROM employees.employees WHERE first_name LIKE ('%ark'); # anything ending in 'ark'
SELECT * FROM employees.employees WHERE first_name LIKE ('ark%'); # anything beginning with 'ark'
SELECT * FROM employees.employees WHERE first_name LIKE ('%ark%'); # anything containing 'ark'
SELECT * FROM employees.employees WHERE first_name NOT LIKE ('%ark%'); # same except excluding this match

# select between ranges (BETWEEN is inclusive)
SELECT * FROM employees.employees WHERE hire_date BETWEEN '1994-12-31' AND '1999-01-01'; # inclusive
SELECT * FROM employees.employees WHERE hire_date BETWEEN '1994-12-31' AND '1999-01-01'; # exclusive

# select NULL or NOT NULL values
SELECT * FROM employees.employees WHERE birth_date IS NULL;
SELECT * FROM employees.employees WHERE birth_date IS NOT NULL;

# conditional operators
SELECT * FROM employees.employees WHERE emp_no != 10010; # not equal to
SELECT * FROM employees.employees WHERE emp_no <> 10010; # not equal to
SELECT * FROM employees.employees WHERE emp_no = 10010; # equal to
SELECT * FROM employees.employees WHERE emp_no < 10010; # less than
SELECT * FROM employees.employees WHERE emp_no <= 10010; # less than or equal to
SELECT * FROM employees.employees WHERE emp_no > 10010; # more than
SELECT * FROM employees.employees WHERE emp_no >= 10010; # more than or equal to

# only extract unique values from column
SELECT DISTINCT gender FROM employees.employees;

# aggregate functions (ignore null values unless told not to)
SELECT COUNT(first_name) FROM employees.employees; # count of non null items
SELECT SUM(salary) FROM employees.salaries; # sum total of column
SELECT MAX(salary) FROM employees.salaries; # max of column
SELECT MIN(salary) FROM employees.salaries; # min of column
SELECT AVG(salary) FROM employees.salaries; # average of column

# count unique occurrences of values in column
SELECT COUNT(DISTINCT first_name) FROM employees.employees;

# ORDER BY a selected field
SELECT * FROM employees.employees ORDER BY first_name; # defaults to ascending
SELECT * FROM employees.employees ORDER BY first_name ASC; # ascending
SELECT * FROM employees.employees ORDER BY first_name DESC; # descending
SELECT * FROM employees.employees ORDER BY first_name, last_name ASC; # multiple column sorting

# GROUP BY to aggregate data
# specify column to group by, must come before order by and after from/where etc.
SELECT first_name, COUNT(first_name) FROM employees.employees GROUP BY first_name ORDER BY first_name ASC;

# add alias to rename column in output
SELECT first_name, COUNT(first_name) AS count_names FROM employees.employees GROUP BY first_name ORDER BY first_name ASC;

# use HAVING instead of WHERE
# HAVING is used on aggregate functions (i.e. when you've used SUM(), AVG(), or GROUP BY etc.)
# whereas WHERE just extracts raw values
SELECT
	emp_no, AVG(salary) AS average_salary
FROM
	employees.salaries
GROUP BY emp_no
HAVING AVG(salary) > 120000
ORDER BY emp_no ASC;

# WHERE is used first to run an initial filter using general functions
# HAVING is then used to further filter using aggregate functions
# note that HAVING cannot use both an aggregate and non-aggregate function at once
# e.g. "HAVING col_name > value AND col_name > SUM(value)" is not allowed
SELECT 
    first_name, COUNT(first_name) AS names_count
FROM
    employees.employees
WHERE
    hire_date > '1999-01-01'
GROUP BY first_name
HAVING COUNT(first_name) < 200
ORDER BY first_name ASC;

# LIMIT lets you crop your output to a specified row count
# useful for top 10s etc.
select emp_no, avg(salary) as average_salary
from salaries
group by emp_no
order by avg(salary) desc
limit 10;