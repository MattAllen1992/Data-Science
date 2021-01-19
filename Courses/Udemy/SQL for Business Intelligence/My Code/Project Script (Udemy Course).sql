# select db to use
use employees_mod;

# manager info query
# get dept, gender, from/to dates and whether or not
# they were active in each calendar year for managers only
select
    d.dept_name,
	ee.gender,
    dm.emp_no,
    dm.from_date,
    dm.to_date,
    e.calendar_year,
    # set active to 1 if role occurred in calendar year
    case
		when e.calendar_year >= year(dm.from_date) and e.calendar_year <= year(dm.to_date) then 1
        else 0
	end as active_emp
# calendar year is defined as the hire date of the manager in each department they were part/manager of
# extract each year from the employees table
from (select year(hire_date) as calendar_year from t_employees group by calendar_year) e
# cross join to 3 other tables to produce cartesian product (i.e. multiply each row of below tables by each year possible)
# below table combo essentially pulls dept info and employee info for managers only
cross join t_dept_manager dm
join t_departments d on d.dept_no = dm.dept_no
join t_employees ee on ee.emp_no = dm.emp_no
order by dm.emp_no, calendar_year;