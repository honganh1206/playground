---
tags:
  - "#study"
cssclasses:
  - center-images
---
### Scan Operator

```sql
CREATE CLUSTERED INDEX IX_ExPlanOperator_ID ON ExPlanOperator (ID)
```

- Although the engine uses clustered index scan, it will still **traverse all the index rows**, similar to the usual Table Scan
- If there is no non-clustered index/the query intends to return most of the data from the table, **the engine will use clustered index scan.**

### Seek Operator

- **A faster way to search** with a `WHERE` condition to **provide the engine with the instructions** to identify the required rows