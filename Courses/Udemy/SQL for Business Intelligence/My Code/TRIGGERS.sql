# triggers let you perform actions before or after specific functions (normally insert, update or delete)
# it allows you to perform checks/cleaning steps etc. depending on particular inputs or changes
# for example, in the below code, anytime a new employee is inserted into the employees table,
# a trigger is run to update the hire date if an invalid (future) date is entered as the start date
drop trigger if exists trig_hire_date;

delimiter $$
create trigger trig_hire_date
before insert on employees
for each row
begin
	if NEW.hire_date > date_format(sysdate(), '%Y-%m-%d') then
		set NEW.hire_date = date_format(sysdate(), '%Y-%m-%d');
	end if;
    if NEW.hire_date <= date_format(sysdate(), '%Y-%m-%d') then
		set NEW.hire_date = OLD.hire_date; # use 'OLD' keyword
	end if;
end$$
delimiter ;

insert into employees values ('999904', '1970-01-31', 'John', 'Johnson', 'M', '2025-01-01');

select * from employees order by emp_no desc;