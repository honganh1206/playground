---
tags:
  - "#study"
cssclasses:
  - center-images
---
- Calculate a new value from the existing row value by performing a scalar computation operation that results a computed value.

```sql
SELECT STD_Name + '_ has achieved _ ' + cast(STD_Grade AS VARCHAR(50)) AS STD_Result
FROM ExPlanOperator_P3 -- Describe the grade of each student
```