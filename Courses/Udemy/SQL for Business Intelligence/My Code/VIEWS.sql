# create a view that implements whatever query you like
# this query creates a view showing latest department each employee has been assigned to
create or replace view v_latest_empl_dept as
select emp_no, dept_no, max(from_date) as from_date, max(to_date) as to_date
from dept_emp
group by emp_no;

# NOTE:
# you can find your query in the db > views > view_name section
# clicking on the "execute query" button (right hand side when hovering over it) shows you its results
# it lets you see the latest, dynamic version of the data based on your query
# you cannot do anything with the view (i.e. edit the underlying data)
# but it allows all users to view the same information without entering the code themselves

# another example of a view showing average salary for all managers
create or replace view v_manager_avg_salary as
select round(avg(s.salary),2) as average_salary
from dept_manager dm
join salaries s
on dm.emp_no = s.emp_no;