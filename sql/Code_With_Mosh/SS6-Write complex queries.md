
### `IN` keyword

```sql
SELECT *
FROM sql_store.products
WHERE product_id NOT IN (
	SELECT DISTINCT product_id
	FROM sql_store.order_items
)
```


> [!tip] Join vs Subqueries
> 
> The decision to use whether `JOIN` or subqueries depends on **performance** and **readability**
> 

### `ALL/ANY` keyword

- `ALL` compares the value with **all** values of a list
- `ANY` compares the value with **each** value in the list and returns all values that satisfy the conditions in the subqueries => Similar to `IN` keyword

```sql
-- ALL
SELECT *
FROM sql_invoicing.invoices
WHERE invoice_total > ALL (
	SELECT invoice_total
    FROM sql_invoicing.invoices
    WHERE client_id = 3
)
-- ANY
SELECT *
FROM sql_invoicing.clients
WHERE client_id = ANY (
	SELECT client_id
	FROM sql_invoicing.invoices
	GROUP BY client_id
	HAVING COUNT(*) >=2 
);
```


### Correlated subqueries

```sql
SELECT *
FROM sql_hr.employees e
WHERE salary > (
	-- Execute this command for EACH employee
	-- Inner query's execution depends on the outer query's row currently being executed
	SELECT AVG(salary)
    FROM sql_hr.employees
    WHERE office_id = e.office_id
)
```


### `EXISTS` operator

```sql
SELECT *
FROM sql_invoicing.clients c
-- Higher performance compared to IN
WHERE EXISTS (
	SELECT client_id
    FROM sql_invoicing.invoices
    WHERE client_id = c.client_id
);
```


### Subqueries in `SELECT` and `FROM` clauses

```sql
-- In SELECT clause
SELECT 
	client_id,
    name,
    -- Inner query's execution depdends on the client_id from the outer query
    (SELECT SUM(invoice_total) FROM sql_invoicing.invoices WHERE client_id = c.client_id) AS total_sales,
    (SELECT AVG(invoice_total) FROM sql_invoicing.invoices) AS average,
    (SELECT total_sales - average)  AS difference
FROM sql_invoicing.clients c

-- In FROM clause
SELECT *
FROM (
	SELECT 
		client_id,
		name,
		-- Inner query's execution depdends on the client_id from the outer query
		(SELECT SUM(invoice_total) FROM sql_invoicing.invoices WHERE client_id = c.client_id) AS total_sales,
		(SELECT AVG(invoice_total) FROM sql_invoicing.invoices) AS average,
		(SELECT total_sales - average)  AS difference
	FROM sql_invoicing.clients c
) AS sales_summary
WHERE total_sales IS NOT NULL
```

---

## Exercises

EP3

```sql
SELECT *
FROM sql_invoicing.clients
WHERE client_id NOT IN (
	SELECT DISTINCT client_id
    FROM sql_invoicing.invoices
);
```

EP4 
```sql
-- Using subqueries
SELECT customer_id, first_name, last_name
FROM sql_store.customers
WHERE customer_id IN (
	SELECT o.customer_id
    FROM sql_store.order_items oi
    JOIN sql_store.orders o USING (order_ID)
    WHERE product_id = 3
);
-- Using joins
SELECT DISTINCT customer_id, first_name, last_name
FROM sql_store.customers
JOIN sql_store.orders o USING (customer_id)
JOIN sql_store.order_items oi USING (order_id)
WHERE oi.product_id = 3
```


EP7

```sql
SELECT *
FROM sql_invoicing.invoices i
WHERE invoice_total > (
	SELECT AVG(invoice_total)
    FROM sql_invoicing.invoices
    WHERE client_id = i.client_id
);
```

EP8

```sql
SELECT *
FROM sql_store.products p
WHERE NOT EXISTS (
	SELECT product_id
    FROM sql_store.order_items
    WHERE product_id = p.product_id
);
```