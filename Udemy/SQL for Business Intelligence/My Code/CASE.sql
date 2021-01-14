# CASE is used almost like an IF statement, where you give different outputs depending on the criteria matched
# it is more flexible than an IF statement because you can test more cases than a simple IF(condition, True, False)
select
	e.emp_no,
	e.first_name,
    e.last_name,
    max(s.salary) - min(s.salary) as salary_difference,
    CASE
		WHEN max(s.salary) - min(s.salary) > 30000 THEN 'Big Increase'
        WHEN max(s.salary) - min(s.salary) BETWEEN 20000 AND 30000 THEN 'Medium Increase'
        WHEN max(s.salary) - min(s.salary) BETWEEN 0 and 20000 THEN 'Small Increase'
        ELSE 'Reduction or No Increase'
	END AS salary_info
from dept_manager dm
join employees e on e.emp_no = dm.emp_no
join salaries s on s.emp_no = dm.emp_no
group by s.emp_no;

# another example where you determine if someone is a manager from another table and then
# adjust your output accordingly
select
	e.emp_no,
    e.first_name,
    e.last_name,
    case
		when dm.emp_no is not null then 'Manager'
		when dm.emp_no is null then 'Employee'
	end as emp_type
from employees e
left join dept_manager dm on e.emp_no = dm.emp_no
where e.emp_no > 109990;

### 2 different ways to do the same thing, the first uses CASE, the second uses IF ###
# emp_no, first and last name (from managers table)
# one col to show diff between max and min salary
# another to say if raise was higher than 30k
# provide more than 1 solution if possible
select
	e.emp_no,
    e.first_name,
    e.last_name,
    max(s.salary) - min(s.salary) as salary_diff,
    case
		when max(s.salary) - min(s.salary) > 30000 then 'Yes'
        else 'No'
	end as large_salary_diff
from dept_manager dm
join employees e on dm.emp_no = e.emp_no
join salaries s on dm.emp_no = s.emp_no
group by dm.emp_no;

select
	e.emp_no,
    e.first_name,
    e.last_name,
    max(s.salary) - min(s.salary) as salary_diff,
    if(max(s.salary) - min(s.salary) > 30000, 'Yes', 'No') as large_salary_diff
from dept_manager dm
join employees e on dm.emp_no = e.emp_no
join salaries s on dm.emp_no = s.emp_no
group by dm.emp_no;

# final example with a group by element at the end and a date comparison
# emp_no, first and last name (from first 100 employees)
# add 'current_empl' saying if they're still employed or not
select
	e.emp_no,
    e.first_name,
    e.last_name,
    CASE
		when max(de.to_date) > sysdate() then 'Active'
        else 'Leaver'
	END as emp_status
from employees e
join dept_emp de
on e.emp_no = de.emp_no
group by de.emp_no
limit 100;