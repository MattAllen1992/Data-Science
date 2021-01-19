# subqueries (or inner queries) are additional queries in brackets (i.e. nested)
# the inner most subqueries are run first, then the next until the outer most query is run last
# it allows you to nest and create really specific queries
# the below query allows you to get first and last name from one table
# when the employee number in that table is present in another table
# the subquery does the check for whether or not values are present in the other table
# subqueries often begin with IN and then brackets
select e.first_name, e.last_name
from employees e
where e.emp_no in (select dm.emp_no from dept_manager dm);

# same as above except with a number of conditionals for a more complex query/subquery
SELECT 
    *
FROM
    dept_manager
WHERE
    emp_no IN (SELECT 
            emp_no
        FROM
            employees
        WHERE
            hire_date > '1990-01-01'
                AND hire_date < '1995-01-01');
			
# EXISTS is similar to in
# it allows you to check if values exist within another table (essentially returns True/False boolean response
# EXISTS works better on large datasets (checks row matches) whilst IN works better on smaller datasets (checks in all values)
# NOTE: it's best practice to order by etc. outside the subquery at the end, to apply sorting to final dataset
select e.first_name, e.last_name
from employees e
where exists (select * from dept_manager dm where e.emp_no = dm.emp_no)
order by e.emp_no;

# alternate example of the above with extra conditions
select *
from employees e
where exists (select *
			  from titles t
              where e.emp_no = t.emp_no
              and t.title = "Assistant Engineer")
order by e.emp_no;

# multiple nested subqueries to complex join multiple tables
# employees 10001 to 10020 and manager 110022
# employees 10021 to 10040 and manager 110039
SELECT 
    A.*
FROM
    (SELECT 
        e.emp_no AS employee_id,
            MIN(de.dept_no) AS department_id,
            (SELECT 
                    emp_no
                FROM
                    dept_manager
                WHERE
                    emp_no = 110022) AS manager_id
    FROM
        employees e
    JOIN dept_emp de ON e.emp_no = de.emp_no
    WHERE
        e.emp_no <= 10020
    GROUP BY e.emp_no
    ORDER BY e.emp_no) AS A 
UNION SELECT 
    B.*
FROM
    (SELECT 
        e.emp_no AS employee_id,
            MIN(de.dept_no) AS department_id,
            (SELECT 
                    emp_no
                FROM
                    dept_manager
                WHERE
                    emp_no = 110039) AS manager_id
    FROM
        employees e
    JOIN dept_emp de ON e.emp_no = de.emp_no
    WHERE
        e.emp_no > 10020
    GROUP BY e.emp_no
    ORDER BY e.emp_no
    LIMIT 20) AS B;
    
# multiple nested subqueries unioned together
# assigning values to a new table
use employees;

drop table if exists emp_manager;

create table emp_manager
(
	emp_no INT(11) NOT NULL,
    dept_no CHAR(4) NULL,
    manager_no INT(11) NOT NULL
);

# A - manager 110022 empl 10001 to 10020
# B - manager 110039 empl 10021 to 10040
# C - manager 110039 manager 110022
# D - manager 110022 manager 110039
insert into emp_manager
SELECT 
    A.*
FROM
    (SELECT 
        e.emp_no,
            MIN(de.dept_no),
            (SELECT 
                    emp_no
                FROM
                    dept_manager
                WHERE
                    emp_no = 110022)
    FROM
        employees e
    JOIN dept_emp de ON e.emp_no = de.emp_no
    WHERE
        e.emp_no <= 10020
    GROUP BY e.emp_no
    ORDER BY e.emp_no) AS A
    UNION
    SELECT 
    B.*
FROM
    (SELECT 
        e.emp_no,
            MIN(de.dept_no),
            (SELECT 
                    emp_no
                FROM
                    dept_manager
                WHERE
                    emp_no = 110039)
    FROM
        employees e
    JOIN dept_emp de ON e.emp_no = de.emp_no
    WHERE
        e.emp_no > 10020
    GROUP BY e.emp_no
    ORDER BY e.emp_no
    LIMIT 20) AS B
UNION
SELECT 
    C.*
FROM
    (SELECT 
        e.emp_no,
            MIN(de.dept_no),
            (SELECT 
                    emp_no
                FROM
                    dept_manager
                WHERE
                    emp_no = 110022)
    FROM
        employees e
    JOIN dept_emp de ON e.emp_no = de.emp_no
    WHERE
        e.emp_no = 110039
    GROUP BY e.emp_no
    ORDER BY e.emp_no) AS C
    UNION
    SELECT 
    D.*
FROM
    (SELECT 
        e.emp_no,
            MIN(de.dept_no),
            (SELECT 
                    emp_no
                FROM
                    dept_manager
                WHERE
                    emp_no = 110039)
    FROM
        employees e
    JOIN dept_emp de ON e.emp_no = de.emp_no
    WHERE
        e.emp_no = 110022
    GROUP BY e.emp_no
    ORDER BY e.emp_no) AS D;
    
    SELECT 
    *
FROM
    emp_manager;