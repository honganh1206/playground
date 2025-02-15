# Indexing

## What are indexes?

- Indexes are **pointers** to records in tables
- Indexes are stored in-memory, instead of data which is stored in disk
- Indexes should be reserved for performance-critical queries

> [!tip]
> 
> Do not index every column. You should design indexes based on your queries, not your tables.

```sql
CREATE INDEX idx_state on customers (state)
```

## Indexing string columns

- To reduce size, we only want to include the prefix of the column

```sql
-- We have to look at our data to determine the number of characters for indexing
CREATE INDEX idx_lastname ON customers (last_name(20))
```

## Full-text indexes

```sql
CREATE FULLTEXT INDEX idx_title_body ON posts(title, body)

SELECT *, MATCH(title, body) AGAINST('react redux')
FROM posts
WHERE MATCH(title, body) AGAINST('react redux')
```


### Relevancy score

- (Only in MySQL) Calculate a decimal number between 0 and 1 for each row containing the searched phrase

### Boolean mode

- Include/exclude certain words 


## Composite Indexes

- Index multiple columns

```sql
CREATE INDEX idx_state_points ON customers (state, points)
```

> [!tip]
> 
> Suggested number of composite indexes is between 4 and 6, but it mostly depends on the data you are working on.


### Order of columns

1. Follow the **leftmost prefix rule**: Supposed you have an index on columns `A,B,C` but your query by filtering on `B` and `C` then the composite index will not be as effective.
2. Put the most frequently used columns first
3. Put the columns with a higher *cardinality* (the number of unique values - search it up with `COUNT(DISTINCT column_name)`) first

## Use indexes for sorting

```sql
-- This only works in MySQL
SHOW STATUS LIKE 'status_name'
```

- Indexed columns can be used with the `ORDER BY` command to reduce cost

> [!warning]
> 
> We can not mix the sorting direction: Supposed that we have two indexes A and B, we can only sort with A or B or A,B or A DESC, B DESC

## Index Maintenance

- Be aware of Duplicate indexes and Redundant Indexes

![[Pasted image 20240821003025.png]]

> [!important]
> 
> Before creating a new index, check existing indexes.


