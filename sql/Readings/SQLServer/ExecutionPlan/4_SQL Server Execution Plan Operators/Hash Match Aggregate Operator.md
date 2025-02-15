---
tags:
  - "#study"
cssclasses:
  - center-images
---
- Steps: 
	1. Build a hash table: Read all rows → Create a hash table
	2. Calculate a hash value based on the group of columns
	3. Scan the table for that hash key + Create new entry in the hash table if not exist
	4. Output one row for each entry in the hash table, each row contains the aggregated value

```sql
SELECT ID
  ,COUNT(*)
FROM ExPlanOperator_P3
GROUP BY ID
```

![[Pasted image 20240802160940.png]]