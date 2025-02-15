
```sql
-- Use database
USE database_name;

SELECT *
FROM customers
WHERE customer_id = 1
ORDER BY first_name
```

- Give an alias to a column

```sql
SELECT (points + 10) * 100 AS 'discount factor'
```

- Select an unique list

```sql
SELECT DISTINCT state FROM customers
```

> [!important]
> 
> `AND` operator has a higher precedence than the `OR` operator


> [!tip]
> 
> The `NOT` operator **negates** other operators
> ```sql
> WHERE NOT (birth_date > '1990-01-01' AND points > 1000)
> -- Is equal to
> WHERE (birth_date < '1990-01-01' OR points <= 1000)
>```


- Combine multiple `OR` conditions with `IN`

```sql
WHERE state IN ('VA', 'FL', 'GA')
```


- Use `BETWEEN` when comparing values

```sql
WHERE points BETWEEN 1000 AND 3000
```

- Use regex with `LIKE`
	- % for *any number of chars*
	- _ for single char

```sql
-- Get last names starting with b case-insensitive
WHERE last_name LIKE 'b%'
-- This query
WHERE last_name LIKE '%field%'
-- is equivalent to
WHERE last_name REGEXP 'field'
```

- Use regex with `REGEXP`
	- `^value` means it must begin with the value,  while `value$` means it must end with the value
	- `|` is like OR for regex
	- `[gim]e` means any value matching with ge/ie/me
	- `[a-h]e` gets all the values matching wh ae/be/ce/.../he

```sql
WHERE last_name REGEXP 'field|mac'
WHERE last_name REGEXP '[gim]e'
```

- Look for records that miss an attribute (null)

```sql
WHERE phone IS NULL
```

- Sorting

```sql
ORDER BY state, first_name DESC
```

- Limit

```sql
-- Skip the first 6 records and get 3 records
SELECT * FROM customers LIMIT 6, 3
```

- Join 3 tables or more 

```sql
SELECT * 
FROM sql_store.orders o 
JOIN sql_store.customers c ON o.customer_id = c.customer_id 
JOIN sql_store.order_statuses os ON o.status = os.order_status_id
```


---

# Exercises

EP1.

```sql
SELECT name, unit_price, (unit_price * 1.1) AS new_price FROM sql_store.products 
```

EP3.

```sql
SELECT * FROM sql_store.order_items WHERE (unit_price * quantity > 30 and order_id = 6)
```

EP7.

```sql
SELECT * from sql_store.customers WHERE address LIKE '%TRAIL%' OR address like '%AVENUE%'
-- OR
SELECT * from sql_store.customers WHERE address REGEXP 'TRAIL' OR address REGEXP 'AVENUE';
SELECT * from sql_store.customers WHERE phone LIKE '%9'
```


EP8.

```sql
SELECT * from sql_store.customers WHERE first_name REGEXP 'elka|ambur';

SELECT * from sql_store.customers WHERE last_name REGEXP 'ey%|on$';

SELECT * from sql_store.customers WHERE last_name REGEXP '^my|se';

SELECT * from sql_store.customers WHERE last_name REGEXP 'b[ru]';
```


EP9.

```sql
SELECT * from sql_store.orders WHERE shipped_date IS null
```

EP10.

```sql
SELECT *, (quantity * unit_price) AS total_price FROM sql_store.order_items WHERE order_id = 2 ORDER BY total_price  DESC
```


EP11.

```sql
SELECT * from sql_store.customers ORDER BY points DESC LIMIT 3
```