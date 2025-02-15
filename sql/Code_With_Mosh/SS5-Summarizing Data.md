
## Aggregate functions

- Take a series of values and return a single aggregated value

```sql
SELECT
	MAX(invoice_total) AS highest,
    MIN(invoice_total) AS lowest,
    AVG(invoice_total) AS average,
    SUM(invoice_total * 1.1) AS total,
    COUNT(client_id) AS total_record
FROM sql_invoicing.invoices
WHERE invoice_date > '2019-07-01'
```

## Grouping

- The `GROUP BY` keyword is used to **group rows that have the same values into summary rows**, typically for use with aggregate functions like `COUNT`, `SUM`, `AVG`, etc.

> [!warning]
> 
> `GROUP BY` must be used before `ORDER BY`


### `HAVING` keyword

- It is similar to the `WHERE` clause but is used specifically with grouped results.

```sql
SELECT
	client_id,
    SUM(invoice_total) AS total_sales
FROM sql_invoicing.invoices
-- Filter before the rows are grouped
WHERE client_id > 1
GROUP BY client_id
-- Filter after the rows are grouped
HAVING total_sales > 500
```


> [!tip] Rule of thumb
> 
> When you have an aggregate function in a select statement, you should group by **ALL** the columns in the `SELECT` statement


### `ROLLUP` keyword 

- Only applies to columns that **aggregate** values (`SUM`)

```sql
SELECT
	state,
    city,
    SUM(invoice_total) AS total_sales
FROM sql_invoicing.invoices i
JOIN sql_invoicing.clients c USING(client_id)
GROUP BY state, city WITH ROLLUP
```

> [!note]
> 
> `ROLLUP` is only available in MySQL




---

## Exercises


### EP1

```sql
SELECT
	'First half of 2019' AS date_range,
    SUM(invoice_total) AS total_sales,
    SUM(payment_total) as total_payments,
    SUM(invoice_total - payment_total) AS what_we_expect
FROM sql_invoicing.invoices
WHERE invoice_date 
	BETWEEN '2019-01-01' AND '2019-06-30'
UNION
SELECT
	'Second half of 2019' AS date_range,
    SUM(invoice_total) AS total_sales,
    SUM(payment_total) as total_payments,
    SUM(invoice_total - payment_total) AS what_we_expect
FROM sql_invoicing.invoices
WHERE invoice_date 
	BETWEEN '2019-07-01' AND '2019-12-31'
UNION
SELECT
	'Total' AS date_range,
    SUM(invoice_total) AS total_sales,
    SUM(payment_total) as total_payments,
    SUM(invoice_total - payment_total) AS what_we_expect
FROM sql_invoicing.invoices
WHERE invoice_date 
	BETWEEN '2019-01-01' AND '2019-12-31'
```

### EP2

```sql
SELECT 
	date,
    pm.name AS payment_method,
    SUM(amount) AS total_payments
FROM sql_invoicing.payments p
JOIN sql_invoicing.payment_methods pm 
	ON p.payment_method = pm.payment_method_id
GROUP BY date, payment_method
ORDER BY date;
```


### EP 3

```sql
SELECT
	c.customer_id,
	c.first_name,
    c.last_name,
    SUM(oi.quantity * oi.unit_price) AS total_sales
FROM sql_store.customers c
JOIN sql_store.orders o USING (customer_id)
JOIN sql_store.order_items oi USING (order_id)
	WHERE c.state = 'VA'
-- Follow the rule iof thumb
GROUP BY c.customer_id,
	c.first_name,
    c.last_name
HAVING total_sales > 100;
```

### EP4

```sql
SELECT 
    pm.name AS payment_method,
    SUM(amount) AS total_payments
FROM sql_invoicing.payments p
JOIN sql_invoicing.payment_methods pm 
	ON p.payment_method = pm.payment_method_id
GROUP BY pm.name WITH ROLLUP;
```