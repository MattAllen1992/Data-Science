# select db to use
use employees_mod;

# average salary by gender
# each year until 2002
# filter by dept
select	
    d.dept_name,
    year(s.from_date) as calendar_year,
    e.gender,
	round(avg(s.salary), 2) as avg_salary
from t_salaries s
join t_dept_emp de on de.emp_no = s.emp_no
join t_departments d on d.dept_no = de.dept_no
join t_employees e on e.emp_no = s.emp_no
group by e.gender, d.dept_name, calendar_year
having calendar_year <= 2002
order by d.dept_no, calendar_year, e.gender;