
## Numeric functions 

### Rounding with `ROUND()/CEILING()/FLOOR()`

```sql
ROUND(5.223, 1) -- Return 5.2
CEILING(5.2) -- Return 6
FLOOR(5.2) -- Return 5
```


## String functions

- `LENGTH()`
- `UPPER()/LOWER()`
- `LTRIM()/RTRIM()/TRIM()`
- `LEFT("ABCDE", 4) -- Return "ABCD"`
- `SUBSTRING("string", 3, 2) -- Return "tr"`
- `LOCATE('n', 'Kinder') -- NOT case-sensitive`
- `CONCAT('first', 'last')`

## Date functions

- `NOW()`
- `YEAR(NOW()) -- Get the year only`
- `EXTRACT(YEAR FROM NOW())`

### Formatting dates and times 

- `DATE_FORNAT(NOW(),'%M %Y')`
- `TIMEFORMAT(NOW(),'%H')`

### Calculate dates and times

- `DATE_ADD(NOW(), INTERVAL 1 DAY)`
- `DATE_SUB(NOW(), INTERVAL 1 DAY)`


## `IFNULL()` and `COALESCE()`


```sql
SELECT
	order_id,
	-- Substitute the null value with something else 
    IFNULL(shipper_id, 'Not assigned') AS shipper,
    -- Return the 1st non-null value in the list
    COALESCE(shipper_id, comments, 'Not assigned') AS shipper_coalesce
FROM sql_store.orders
```


## `CASE`

```sql
SELECT
	order_id,
    order_date,
    CASE
		WHEN YEAR(order_date) = YEAR(NOW()) THEN 'Active'
        WHEN YEAR(order_date) = YEAR(NOW()) - 5 THEN 'Last Year'
        WHEN YEAR(order_date) < YEAR(NOW()) - 5 THEN 'Archived'
        ELSE 'Future'
	END AS category
FROM sql_store.orders;
```

—

## Exercises

EP6

```sql
SELECT
	CONCAT(first_name, ' ', last_name) AS customer,
    IFNULL(phone, 'Unknown') AS phone
FROM sql_store.customers;
```

EP7

```sql
SELECT 
    p.product_id,
    p.name,
    COUNT(*) AS orders,
    IF (
		COUNT(*) > 1,
        'Once',
        'Many times'
    ) AS frequency
FROM sql_store.products p
JOIN sql_store.order_items oi
USING (product_id)
GROUP BY p.product_id, p.name;
```

EP8

```sql
SELECT
	CONCAT(first_name, ' ', last_name) AS customer,
    points,
    CASE
		WHEN points > 3000 THEN 'Gold'
        WHEN 2000 < points < 3000 THEN 'Silver'
        ELSE 'Bronze'
	END AS category
FROM sql_store.customers;
```