---
id: Order of columns
aliases: []
tags: []
---

# Order of columns in SQL

[Ref](https://stackoverflow.com/questions/2292662/how-important-is-the-order-of-columns-in-indexes) 

TL,DR: Narrow down  the number of results to deal with before the next step for performance gains

```text
Cols
  1   2   3
-------------
|   | 1 |   |
| A |---|   |
|   | 2 |   |
|---|---|   |
|   |   |   |
|   | 1 | 9 |
| B |   |   |
|   |---|   |
|   | 2 |   |
|   |---|   |
|   | 3 |   |
|---|---|   |
```

Starting with column 1 is more efficient, since doing so allows us to *eliminate half of the data*. From that point any subsequent filtering on Column 2 only needs to work on a smaller subset 

When querying on column 3, the database optimizer will likely ignore the index since **it cannot use the hierarchical structure**. For that reason we must **scan the entire index** to find matching values  
