### STORED ROUTINES ###
# There are 2 types of stored routines: procedures and functions
# Functions return a value, procedures don't
# Functions must be 'selected', procedures must be 'called'
# Procedures can output multiple values (not technically return) whilst functions only output one
# If you're making changes to your DB (e.g. insert, update, delete) then you should use a procedure without an 'out' parameter
# 	this is because functions must return a value, and DB updates don't necessarily do this
# Because you can 'select' functions, you can use their aliases within other select functions

### 1) PROCEDURES ###
# drop exiting procedure/routine if it exists
DROP procedure IF EXISTS select_employees;

# define a new delimiter to use, otherwise you'll encounter issues when calling the procedure
DELIMITER $$
CREATE procedure select_employees() # define the name of the routine
BEGIN # always begin with BEGIN and end with END
	SELECT * FROM employees # this query can be anything you want it to be
    LIMIT 10;
END $$ # end your routine
DELIMITER ; # reset the delimiter back to its original state

# invoke/call/reference your routine to produce the relevant results
# note that you don't have to specify db name or use parentheses if there are no parameters
call select_employees();
call select_employees;
call employees.select_employees();
call employees.select_employees;

# drop your procedure from the DB
drop procedure select_employees();

# NOTE
# you can also create routines/procedures using the schemas window
# right-click on your procedure and create new
# or click on the wrench symbol to edit existing
# the lightning bolt will run it and produce the desired output

# stored routine/procedure with an input parameter
# just like creating a python method that takes a specific input value
DELIMITER $$
CREATE procedure empl_sal(IN p_emp_no INTEGER)
BEGIN
	SELECT e.first_name, e.last_name, s.salary, s.from_date, s.to_date
    FROM employees e
    JOIN salaries s
    ON e.emp_no = s.emp_no
    WHERE e.emp_no = p_emp_no;
END $$
DELIMITER ;

call empl_sal(11300);

# specify an output parameter as well
# this lets you insert the output of your query into a specific variable
DELIMITER $$
CREATE procedure empl_avg_sal_out(IN p_emp_no INTEGER, OUT p_avg_sal DECIMAL(10,2))
BEGIN
	SELECT avg(s.salary)
    INTO p_avg_sal
    FROM employees e
    JOIN salaries s
    ON e.emp_no = s.emp_no
    WHERE e.emp_no = p_emp_no;
END$$
DELIMITER ;

call empl_avg_sal_out(11300, @p_avg_sal);

### ASSIGN OUTPUT VALUES FROM PROCEDURE TO VARIABLE ###
# create variable and instantiate with 0
SET @v_emp_no = 0;

# prodedure to get employee number given first and last name
DELIMITER $$
CREATE procedure emp_info(IN p_first_name VARCHAR(255), IN p_last_name VARCHAR(255), OUT p_emp_no INTEGER)
BEGIN
	SELECT MAX(e.emp_no)
    INTO p_emp_no
	FROM employees e
	WHERE e.first_name = p_first_name AND e.last_name = p_last_name;
END$$

DELIMITER ;

# call procedure and assign output to variable
CALL emp_info('Georgi', 'Facello', @v_emp_no);

# select output variable to show value
SELECT @v_emp_no;

### 2) FUNCTIONS ###
# Functions differ from procedures in that:
# a) You don't need to specify IN or OUT, instead you put inputs/arguments in the brackets and simply specify data type of out variables
# b) you must declare your output variable and its type (which must match what's stated above it)
# c) you must return your output variable at the end
# d) you cannot call a function, you must select it to use it
# Note that the DETERMINISTIC keyword is used to prevent error 1418, which essentially occurs because SQL thinks you're trying
# to overwrite underlying data in your query, so you must tell it explicitly that you're not
DELIMITER $$
create function f_avg_salary(p_emp_no INTEGER) returns decimal(10,2)
DETERMINISTIC
BEGIN
	declare v_avg_salary decimal(10,2);
    
SELECT 
    AVG(s.salary)
INTO v_avg_salary FROM
    employees e
        JOIN
    salaries s ON e.emp_no = s.emp_no
WHERE
    e.emp_no = p_emp_no;
    
    return v_avg_salary;
END$$
DELIMITER ;

select f_avg_salary(11300);

# you can use functions as arguments/parameters in other queries
# for example, below, we use the average salary function to return the averge salary for a specific employee
# after setting a variable to contain a specific value
set @v_emp_no = 11300;
select
	emp_no,
	first_name,
	last_name,
    f_avg_salary(@v_emp_no) as avg_salary
from employees
where emp_no = @v_emp_no;