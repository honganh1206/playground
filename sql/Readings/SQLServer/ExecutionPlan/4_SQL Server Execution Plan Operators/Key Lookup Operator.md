---
tags:
  - "#study"
cssclasses:
  - center-images
---
- Key Lookup = Clustered RID Lookup
- With the clustered index, the engine will pointers (returned from non-clustered index) will point to the clustered index

![[Pasted image 20240801164647.png|The nested loops now joins the data from the Index seek + Key lookup as the engine is not able to retrieve rows in a shot]]

> [!warning]
> 
> For each row found by the Index seek, a separate key lookup using pointers is performed to fetch the rest of the data => The SQL Server Engine cannot retrieve all the rows “in one shot” because it has to perform these separate lookups for each row.


- The Key Lookup is also **resource-consuming** as the RID lookup

```sql
CREATE INDEX IX_ExPlanOperator_P2_EmpFirst_Name ON ExPlanOperator_P2 (EmpFirst_Name) INCLUDE (ID,EmpLast_name, EmpAddress, EmpPhoneNum) WITH (DROP_EXISTING = ON)
```