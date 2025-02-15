## Inserting


- Insert one row/multiple rows

```sql
-- Insert 1 row
INSERT INTO sql_store.customers (
	first_name, 
	last_name,
    birth_date,
    address,
    city,
    state)
VALUES ( 
	'John', 
    'Smith', 
    "2000-01-01",
    'address',
    'city',
    'CA')
```

- Insert hierarchical data

```sql
INSERT INTO sql_store.orders (customer_id, order_date, status)
VALUES (1, '2019-01-03', 1);

INSERT INTO sql_store.order_items
VALUES 
	(LAST_INSERT_ID(), 1, 1, 2.95),
    (LAST_INSERT_ID(), 2, 1, 3.95);
```

- Create a copy of a table

```sql
CREATE TABLE sql_store.orders_archived_1 AS
SELECT * FROM sql_store.orders
-- Copy data to another table
INSERT INTO sql_store.orders_archived_1
SELECT *
FROM sql_store.orders
WHERE order_date < '2019-01-01'
```

## Updating

### Using subqueries

```sql
UPDATE sql_invoicing.invoices
SET payment_total = invoice_total * 0.6, payment_date = due_date
WHERE client_id IN 
	(SELECT client_id 
	FROM sql_invoicing.clients 
	WHERE state IN ('CA', 'NY'));
```


## Deleting


```sql
DELETE FROM sql_invoicing.invoices
WHERE client_id = (
	SELECT client_id
    FROM sql_invoicing.clients
    WHERE name = 'Myworks'
);
```


## Restore the database

---

## Exercises

- EP5

```sql
CREATE TABLE sql_invoicing.invoicing_archived AS
SELECT 
		i.invoice_id,
        i.number,
        c.name AS client,
        i.invoice_total,
        i.payment_total,
        i.invoice_date,
        i.payment_date,
        i.due_date
FROM sql_invoicing.invoices i
JOIN sql_invoicing.clients c
	USING(client_id)
WHERE i.payment_date IS NOT NULL
```


- EP8

```sql
UPDATE sql_store.orders
SET comments = "Gold Customer"
WHERE customer_id in 
	(SELECT customer_id 
	FROM sql_store.customers 
	WHERE points > 3000);
```