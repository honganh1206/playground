---
tags:
  - "#study"
cssclasses:
  - center-images
---

```sql
CREATE INDEX IX_ExPlanOperator_FirstName on ExPlanOperator(First_Name)
```

- If a non-clustered index is provided + the query intends to return only a small number of records => Non-clustered index will be prioritized by the optimizer

![[Pasted image 20240801143958.png|Seeking specific rows only with non-clustered index]]