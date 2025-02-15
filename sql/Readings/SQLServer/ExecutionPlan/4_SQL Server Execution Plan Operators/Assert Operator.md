---
tags:
  - "#study"
cssclasses:
  - center-images
---
- Verify if the inserted values meet the previously defined CHECK or FOREIGN KEY constraints on the table

```sql

ALTER TABLE ExPlanOperator_P3 ADD CONSTRAINT CK_Grade_Positive CHECK (STD_Grade >0) -- Add a constraint


INSERT INTO ExPlanOperator_P3 VALUES ('GG','1998-01-28','HH',74) -- The query meets the constraint
```

![[Pasted image 20240802135224.png]]