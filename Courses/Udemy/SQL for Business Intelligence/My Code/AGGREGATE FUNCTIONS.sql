# a.k.a summarizing functions
COUNT(), COUNT(DISTINCT col_name) # non-null count of (distinct) occurrences
COUNT(*) # count of occurrences, includes nulls
SUM() # sums values
MIN() # min value
MAX() # max value
AVG() # average value

# use these with SELECT
select count(distinct dept_no)
from deptartments;

# and with WHERE
select sum(salary)
from salaries
where emp_no > 100000;

# use ROUND with numerical arguments to round decimals to specific number of digits
select emp_no, round(avg(salary), 2) as average_salary
from salaries;

# use IFNULL (essentially IFERROR) to return a value if found
# or a specific alternate value if it's null
select IFNULL(dept_name, 'N/A')
from dept_emp;

# use COALESCE if you have more than 2 fields
# it will try the first value, then the second, then the third (or more)
# until it finds a non-null value
select COALESCE(dept_manager, dept_name, 'N/A')
from dept_emp;

# you can use COALESCE with one value and it'll create
# a single column with just this value in it, it's useful for placeholders during design
select COALESCE('fake value') as fake_column
from dept_emp;