# delete specific rows from tables
# if you omit the WHERE clause, it'll delete all rows in the table
# if you delete a record which is linked via keys to other tables, you'll lose the linked records too
delete from departments
where dept_no = 'd009';

# DROP
# this drops everything (table, records and relationships/constraints)

# TRUNCATE
# this drops all rows and resets indexes
# table and relationships/constraints remain

# DELETE
# removes specific rows depending on conditions in WHERE
# doesn't reset indexes (i.e. index will continue from n)