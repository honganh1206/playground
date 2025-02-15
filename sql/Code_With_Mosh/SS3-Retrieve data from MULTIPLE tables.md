
- Retrieve data from another table with `JOIN` with conditions `ON`

```sql
SELECT * FROM sql_store.orders JOIN sql_store.customers ON orders.customer_id = customers.customer_id
```

- Another way to make aliases

```sql
FROM orders o JOIN customers c ON o.customer_id = customers.customer_id
```

- Retrieve data from the same table

```sql
SELECT e.employee_id, e.first_name, m.first_name as manager FROM sql_hr.employees e JOIN sql_hr.employees m ON e.reports_to = m.employee_id;
```


# Composite primary key

- CPK contains **more than 1 column** => The combination of values in those columns must be **unique** for each record in the table.

```css
| StudentID | CourseID | Name      | Grade |
|-----------|----------|-----------|-------|
| 1         | 101      | John      | A     |
| 2         | 102      | Jane      | B     |
| 3         | 101      | Bob       | C     |
| 1         | 101      | John      | A     |   <-- Not allowed due to the composite primary key

```


# Join

- Compound join statements 

```sql
SELECT *
FROM sql_store.order_items oi
JOIN sql_store.order_item_notes oin
-- Compound join statements
	ON oi.order_id = oin.order_Id
	AND oi.product_id = oin.product_id
```

## Outer joins

### Left/Right join

- Return elements from the **left/right table** whether the `ON` condition is true or not

```sql
SELECT *
-- Data from orders (in case of left) or customers (in case of right) will be returned
-- Whether the ON condition is true or not
FROM sql_store.orders o
LEFT JOIN sql_store.customers c
	ON c.customer_id = o.customer_id
ORDER BY c.customer_id

-- Joining multiple tables
SELECT c.customer_id, c.first_name, o.order_id, sh.name AS shipper_name
FROM sql_store.customers c
JOIN sql_store.orders o
	ON c.customer_id = o.customer_id
LEFT JOIN sql_store.shippers sh
	ON o.shipper_id = sh.shipper_id
ORDER BY c.customer_id;
```

> [!tip]
> 
> **Avoid right joins and use left joins** instead because you can easily read the 'from tableA' as tableA being (visually) the 'primary' table


### Self Join

```sql
SELECT 
	e.employee_id,
    e.first_name,
    m.first_name as manager
FROM sql_hr.employees e
-- Left join to get the CEO
LEFT JOIN sql_hr.employees m
	ON e.reports_to = m.employee_id
```


### Natural Join

- Join columns sharing the same name from the two tables

```sql
SELECT 
	o.order_id,
    c.first_name
FROM sql_store.orders o
NATURAL JOIN sql_store.customers c
```

### Cross Join as an Implicit syntax

- Join **every** record from the 1st table with **every** record from the 2nd table

```sql
SELECT 
	c.first_name AS customer,
    p.name AS product
FROM sql_store.customers c
CROSS JOIN sql_store.products p
ORDER BY c.first_name
```

# `USING` clause

```sql
SELECT p.product_id, p.name, oi.quantity
FROM sql_store.order_items oi
RIGHT JOIN sql_store.products p
-- Equals to oi.product_id = p.product_id
USING(product_id);
-- Multiple usings
SELECT *
FROM sql_store.order_items oi
JOIN sql_store.order_item_notes oin
	USING(order_id, product_id)
```


# `UNIONS`

- **Combine** records from multiple queries

```sql
SELECT c.first_name
FROM sql_store.customers c
UNION
SELECT s.name
FROM sql_store.shippers s
```

---

# Exercises

EP1.

```sql
SELECT order_id, oi.product_id, quantity, oi.unit_price FROM sql_store.order_items oi JOIN sql_store.products p ON oi.product_id = p.product_id
```

EP4.

```sql
SELECT p.date, p.invoice_id, p.amount, c.name as customer_name, pm.name as payment_method_name
FROM sql_invoicing.payments p
JOIN sql_invoicing.payment_methods pm ON p.payment_method = pm.payment_method_id
JOIN sql_invoicing.clients c ON p.client_id = c.client_id
```


EP7.

```sql
SELECT p.product_id, p.name, oi.quantity
FROM sql_store.order_items oi
RIGHT JOIN sql_store.products p
	ON oi.product_id = p.product_id
```


EP8.

```sql
SELECT
    o.order_date,
    o.order_id, 
	c.first_name, 
    sh.name AS shipper_name,
    os.name as status
FROM sql_store.customers c
JOIN sql_store.orders o
	ON c.customer_id = o.customer_id
LEFT JOIN sql_store.shippers sh
	ON o.shipper_id = sh.shipper_id
LEFT JOIN sql_store.order_statuses os
	ON o.status = os.order_status_id
ORDER BY os.name;
```

EP10.

```sql
SELECT 
	p.date,
    c.name AS clients,
    p.amount,
    pm.name AS payment_method
FROM sql_invoicing.payments p
LEFT JOIN sql_invoicing.clients c
	USING(client_id)
LEFT JOIN sql_invoicing.payment_methods pm
	ON p.payment_method = pm.payment_method_id
```

EP13. 

```sql
SELECT c.customer_id, c.first_name, c.points, 'Bronze' AS type
FROM sql_store.customers c
WHERE c.points BETWEEN 2000 AND 3000
UNION
SELECT c.customer_id, c.first_name, c.points, 'Silver' AS type
FROM sql_store.customers c
WHERE c.points BETWEEN 2000 AND 3000
UNION
SELECT c.customer_id, c.first_name, c.points, 'Gold' AS type
FROM sql_store.customers c
WHERE c.points > 3000
```