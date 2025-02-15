---
tags:
  - "#study"
cssclasses:
  - center-images
---
```sql
CREATE INDEX IX_ExPlanOperator_P2_EmpFirst_Name ON ExPlanOperator_P2 (EmpFirst_Name)

SELECT * FROM ExPlanOperator_P2 WHERE EmpFirst_Name = 'BB' -- This is executed without a clustered index
```

- Without a clustered index, the query above only returns a list with employee’s first name + pointers to where the rest of the data we want to retrieve => Heap table that still requires a full table scan

![[Pasted image 20240801163235.png|High I/O overhead as two different operations were performed]]