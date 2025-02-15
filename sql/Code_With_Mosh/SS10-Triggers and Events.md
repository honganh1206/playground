## Triggers

> [!info] Definition
> 
> A block of SQL code that **automatically gets executed** before/after an INSERT/UPDATE/DELETE statement

- AFTER INSERT trigger

```sql
DELIMITER $$ -- Trigger itself contains semicolons => Avoid confusion between semicolons for SQL statement vs. those to terminate statements
CREATE TRIGGER payments_after_insert
	AFTER INSERT ON sql_invoicing.payments
    FOR EACH ROW -- Trigger one for each row inserted
BEGIN
	UPDATE sql_invoicing.invoices
    SET payment_total = payment_total + NEW.amount -- Increment the payment_total by the amount of the inserted row
    WHERE invoice_id = NEW.invoice_id;
END $$

DELIMITER ;
```

- AFTER DELETE trigger

```sql
DELIMITER $$

DROP TRIGGER IF EXISTS sql_invoicing.payments_after_DELETE;

CREATE TRIGGER sql_invoicing.payments_after_DELETE
	AFTER DELETE ON sql_invoicing.payments
    FOR EACH ROW
BEGIN
	UPDATE sql_invoicing.invoices
    SET payment_total = payment_total - OLD.amount
    WHERE invoice_id = OLD.invoice_id;
    
    INSERT INTO sql_invoicing.payments_audit
    VALUES (OLD.client_id, OLD.date, OLD.amount, 'Delete', NOW());
END $$

DELIMITER ;
```

## Events


> [!info] Definition
> 
> A task/block of SQL code that gets executed according to a schedule


```sql
DELIMITER $$

CREATE EVENT sql_invoicing.yearly_delete_stale_audit_rows
ON SCHEDULE
	EVERY 1 YEAR STARTS '2024-05-04' ENDS '2026-05-04'
DO BEGIN
	DELETE FROM sql_invoicing.payments_audit
    WHERE action_date < NOW() - INTERVAL 1 YEAR;
END $$

DELIMITER ;

-- Edit the events

ALTER EVENT sql_invoicing.yearly_delete_stale_audit_rows ENABLE;
```