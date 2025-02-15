
> [!info] Definition
> 
> Stored procedures are precompiled SQL code that can be executed later. They are useful for encapsulating frequently used or complex SQL logic.


## Benefits

- Store and organize SQL`
- Faster execution
- Data security


```sql
-- Make all statements as one unit
DELIMITER $$ -- Change delimiter to $$ (In MySQL only)
CREATE PROCEDURE sql_invoicing.get_clients()
BEGIN
	SELECT * FROM sql_invoicing.clients;
END$$

DELIMITER ;
```


> [!warning] Why we need delimiters
> 
> As each statement ends with a semicolon, using the semicolon for the entire script (Procedures/Triggers/etc.,) would lead to confusion.


```sql
-- Original script with semicolon as delimiter
CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(50)
);

INSERT INTO users (id, name) VALUES (1, 'John');
INSERT INTO users (id, name) VALUES (2, 'Alice');

-- Trigger definition
CREATE TRIGGER my_trigger
AFTER INSERT ON users
FOR EACH ROW
BEGIN
    INSERT INTO audit_table (user_id, action) VALUES (NEW.id, 'insert'); -- This would lead to confusion as a statement terminator
END;

```


## Parameters


```sql
DELIMITER $$
CREATE PROCEDURE sql_invoicing.get_clients_by_state
(
	p_state CHAR(2) -- Parameter
)
BEGIN
	SELECT *
	FROM sql_invoicing.clients c
    WHERE c.state = p_state;
END$$

DELIMITER $$
CREATE PROCEDURE sql_invoicing.get_clients_by_state
(
	state CHAR(2) -- Parameter
)
BEGIN
	IF state IS NULL OR state = '' THEN
		SET state = 'CA'; -- Default parameter
	END IF;
    
	SELECT *
	FROM sql_invoicing.clients c
    WHERE c.state = state;
END$$
DELIMITER ;
```

### Parameter validation

```sql
CREATE DEFINER=`root`@`localhost` PROCEDURE `make_payment`(
	invoice_id INT,
    payment_amount DECIMAL(9, 2), -- Total num of digits, num of digits after decimal point
    payment_date DATE)
BEGIN
	-- Validate payment => Throw exception
    IF payment_amount <= 0 THEN
		SIGNAL SQLSTATE '22003' -- Numeric value out of range
			SET MESSAGE_TEXT = 'Invalid payment amount';
	END IF;
-- Rest of the procedure
```

- [Link to IBM for reference](https://www.ibm.com/docs/en/db2-for-zos/13?topic=codes-sqlstate-values-common-error)

### Output parameters

```sql
CREATE DEFINER=`root`@`localhost` PROCEDURE `get_unpaid_invoices_for_client`(
	client_id INT,
    OUT invoices_count INT,
    OUT invoices_total DECIMAL(9, 2)
)
BEGIN
	SELECT COUNT(*), SUM(invoice_total)
    INTO invoices_count, invoices_total -- Output params 
    FROM invoices i
    WHERE i.client_id = client_id AND payment_total = 0;
END
```


## Variables 

```sql
CREATE DEFINER=`root`@`localhost` PROCEDURE `get_risk_factor`()
BEGIN
-- risk factor = invoice total / invoice count * 5
	DECLARE risk_factor DECIMAL(9,2) DEFAULT 0; -- Local variables inside procedures only
    DECLARE invoices_total DECIMAL(9,2);
    DECLARE invoices_count INT;
    
    SELECT COUNT(*), SUM(invoice_total)
    INTO invoices_count, invoices_total
    FROM invoices;
    
    SET risk_factor = invoices_total / invoices_count * 5;
    
    SELECT risk_factor;
END

```

## Functions 

```sql
CREATE DEFINER=`root`@`localhost` FUNCTION `get_risk_factor_for_client`(
	client_id INT
) RETURNS int
    READS SQL DATA
BEGIN
	DECLARE risk_factor DECIMAL(9,2) DEFAULT 0; -- Local variables inside procedures only
    DECLARE invoices_total DECIMAL(9,2);
    DECLARE invoices_count INT;
    
SELECT 
    COUNT(*), SUM(invoice_total)
INTO invoices_count , invoices_total FROM
    invoices i
WHERE
    i.client_id = client_id;
    
    SET risk_factor = invoices_total / invoices_count * 5;
	RETURN IFNULL(risk_factor, 0);
END

-- Drop the function
DROP FUNCTION IF EXISTS function_name;
```


## Other conventions

> [!tip]
> 
> Follow the team’s naming convention, do not reinvent the wheel.

—

## Exercises


EP2

```sql
DELIMITER $$
CREATE PROCEDURE sql_invoicing.get_invoices_with_balance()
BEGIN
	SELECT *
	FROM sql_invoicing.invoices
    WHERE invoice_total - payment_total > 0;
END$$

DELIMITER ;
```

EP6

```sql
DELIMITER $$
CREATE PROCEDURE sql_invoicing.get_payment
(
	client_id INT,
    payment_method_id TINYINT
)
BEGIN
	SELECT *
	FROM sql_invoicing.payments p
    WHERE p.client_id = IFNULL(client_id, p.client_id)
		AND p.payment_method = IFNULL(payment_method_id, p.payment_method);
END$$

DELIMITER ;
```