# create sales DB
CREATE DATABASE IF NOT EXISTS sales;

# specify sales db as the one to execute following code on
USE sales;

# create sales table
# auto_increment creates unique values from 1 onwards for each row
# not null ensures valid values are present
CREATE TABLE sales
(
	purchase_number INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    date_of_purchase DATE NOT NULL,
    customer_id INT,
    item_code VARCHAR(10) NOT NULL
);

# alter existing table
ALTER TABLE customers

# add new column, enumerated to only allow m or f, after last_name col
ADD COLUMN gender ENUM('M', 'F') AFTER last_name;

# add specific values into row, specifying column names
INSERT INTO customers (first_name, last_name, gender, email_address, number_of_complaints)
VALUES ('John', 'Mackinley', 'M', 'john.mckinley@365careers.com', 0)
;

# add and remove not null constraint
ALTER TABLE sales
MODIFY COLUMN phone_number VARCHAR(255) NULL;

ALTER TABLE sales
CHANGE COLUMN phone_number phone_number VARCHAR(255) NOT NULL;