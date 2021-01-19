# update specific record(s) based on criteria
# don't have to update all columns (especially if record already exists)
# if you don't provide WHERE, all rows will be updated
update employees
set birth_date = '1980-01-01', first_name = 'Mork', last_name = 'Porkens', gender = 'M'
where emp_no = 999903;

# commit current changes to db (e.g. new records etc.)
# make a change which accidentally overwrites all values in a table
# use rollback to revert back to the state of the last commit
# all changes have been reverted
# NOTE: you can only rollback to the previous commit, no further
commit;

update departments_dup
set dept_no = 'd011', dept_name = 'Test';

select *
from departments_dup;

rollback;