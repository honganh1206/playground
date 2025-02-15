

## Benefits of using views

- Simplify queries
- Reduce the impact of changes
- Restrict access to the data

## Creating views

- Views are similar to **virtual tables**
- Views do not store values

```sql

CREATE VIEW sql_invoicing.sales_by_client AS
SELECT
	c.client_id,
    c.name,
    SUM(invoice_total) AS total_sales
FROM sql_invoicing.clients c
JOIN sql_invoicing.invoices i USING (client_id)
GROUP BY client_id, name;
-- Check in the Views section
```


## Altering/Dropping/Updating views

```sql
-- Change data in a view
DELETE FROM view_name
WHERE condition

WITH CHECK OPTION; -- Prevent data from views from being altered
```

## Exercises

EP1

```sql
CREATE VIEW sql_invoicing.clients_balance AS
SELECT
	c.client_id,
    c.name,
    SUM(invoice_total - payment_total) AS balance
FROM sql_invoicing.clients c
JOIN sql_invoicing.invoices USING (client_id)
GROUP BY client_id, name;
```