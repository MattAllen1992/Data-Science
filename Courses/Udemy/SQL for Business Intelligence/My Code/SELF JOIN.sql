# you can join a table to itself and use multiple aliases to specify each version
SELECT 
    e1.emp_no, e2.dept_no, e2.manager_no
FROM
    emp_manager e1
        JOIN
    emp_manager e2 ON e1.emp_no = e2.emp_no
WHERE
    e1.emp_no IN (SELECT 
            manager_no
        FROM
            emp_manager)
;