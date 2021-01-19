# INNER JOIN (shared central point of a venn diagram)
# use aliases for each table (e.g. t1 and t2 right after table names)
# define columns you want to end up with using dot notation
# select the table you're joining to (i.e. left table using FROM)
# and the table you're joining from (i.e. right table using JOIN)
# any nulls will be ommitted from the output, only joins on matching values
# you can write inner join or join, they are synonymous
select m.dept_no, m.emp_no, d.dept_name
from dept_manager_dup m
inner join departments_dup d on m.dept_no = d.dept_no # these names don't have to be the same
where dept_no > 5 # optional: add conditional criteria to join
group by m.dept_no # this line ensures you remove duplicates from dept_no
order by m.dept_no asc;

# LEFT JOIN (left and middle of venn diagram)
# left table is on the FROM line
# right table is on the LEFT JOIN line
# it keeps all rows from left table and matching values from right table
select m.dept_no, m.emp_no, d.dept_name
from dept_manager_dup m
left join departments_dup d on m.dept_no = d.dept_no
group by m.dept_no
order by m.dept_no asc;

# LEFT OUTER JOIN (identical to the above, interchangeable)
select m.dept_no, m.emp_no, d.dept_name
from dept_manager_dup m
left outer join departments_dup d on m.dept_no = d.dept_no
group by m.dept_no
order by m.dept_no asc;

# RIGHT JOIN (right and middle of venn diagram)
# same as left join but reverse table direction
select m.dept_no, m.emp_no, d.dept_name
from dept_manager_dup m
right join departments_dup d on m.dept_no = d.dept_no
group by m.dept_no
order by m.dept_no asc;

# the following 2 blocks of code produce identical outputs
# the first is the old method of using WHERE instead of JOIN to perform an inner join
# both work but WHERE is old, slower and less flexible (you cannot use right join, outer join etc.)
# the WHERE method is not redundant though, it can be used with EXISTS for niche operations (more later)
select d.dept_no, d.dept_name, dm.emp_no
from departments_dup d, dept_manager_dup dm
where d.dept_no = dm.dept_no;

select d.dept_no, d.dept_name, dm.emp_no
from departments_dup d
join dept_manager_dup dm on d.dept_no = dm.dept_no;

# JOIN and WHERE for conditional joins
select e.emp_no, e.first_name, e.last_name, e.hire_date, t.title, t.from_date, t.to_date
from employees e
join titles t on e.emp_no = t.emp_no
where e.first_name = 'Margareta' AND e.last_name = 'Markovitch'
order by emp_no;

# CROSS JOIN (cartesian product i.e. multiply all values in one table by all values in the other)
# all possible combinations of rows in table 1 with rows in table 2
# e.g. show all departments a manager can work for
select dm.*, d.*
from departments d
cross join dept_manager_dup dm
where d.dept_no = 'd009' # optional: conditional filtering
order by d.dept_no asc;

# you can use aggregate functions (e.g. group by)
# in combination with joins
select e.gender, avg(s.salary) as average_salary
from salaries s
join employees e on s.emp_no = e.emp_no
group by e.gender;

# you can run multiple consecutive joins to join multiple tables at once
# just ensure you have a clear plan of what you're pulling from where
# it's really easy to pull in incorrect, duplicate etc. info
select e.first_name, e.last_name, e.hire_date, t.title, t.from_date, t.to_date, d.dept_name
from dept_manager dm
join employees e on dm.emp_no = e.emp_no
join titles t on e.emp_no = t.emp_no
join departments d on dm.dept_no = d.dept_no
where t.title = 'Manager'
order by dm.emp_no;

# complex join with aliases, group by/having etc.
select d.dept_name, avg(s.salary) as average_salary
from departments d
join dept_manager dm on d.dept_no = dm.dept_no
join salaries s on dm.emp_no = s.emp_no
group by d.dept_name
having average_salary > 60000
order by average_salary desc;

# you can use UNION (no duplicates) or UNION ALL (incl. duplicates)
# this essentially appends one table to the next
# you must specify columns with the NULL AS prefix if they don't 
select col_1, NULL AS col_2
from table_1 t1
union select NULL AS col_1, col_2
from table_2 t2;