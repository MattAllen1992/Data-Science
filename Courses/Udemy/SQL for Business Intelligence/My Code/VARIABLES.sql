# There are 3 types of variables: local, session and global

### LOCAL VARIABLES ###
# Local variables are created within BEGIN END blocks
# In the below code, v_avg_salary is a local variable
# It cannot be accessed outside of this function
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

### SESSION VARIABLES ###
# A session begins when you've made a connection (e.g. the one I'm currently in)
# And ends when you close down and log off this connection
# Session variables are therefore variables which are only accessible during the current session
# And will terminate/disappear once the current session has closed
# You will be the only one who can see the variable in your session, noone else using the DB will be able to
SET @s_var1 = 3;
SELECT @s_var1;

### GLOBAL VARIABLES ###
# These can be accessed by any user on any connection providing the same server is being used
# You can't just set any variable as a global variable, you must use system variables
# e.g. max_connections = 1000 (limit the number of sessions allowed on a server)
SET GLOBAL max_connections = 10;
SET @@global.var_name = value; # 2 ways of writing the same thing

# You can distinguish local variables from system variables because:
# Only system variables can be set as global
# Only user created variables can be set locally
# Session variables can be both user defined and global but global variables often stop you using SESSION to create them