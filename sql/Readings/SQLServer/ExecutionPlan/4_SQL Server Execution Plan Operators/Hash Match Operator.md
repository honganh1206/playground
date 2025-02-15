---
tags:
  - "#study"
cssclasses:
  - center-images
---
- When joining two tables together, the engine divides the data into equally sized buckets (Hashing table) using an algorithm for data distribution (Hashing function)

```sql
-- Create the table
CREATE TABLE ExPlanOperator_JOIN (
  STD_ID INT,
  STD_AbsenceDays INT
);
GO

-- Initialize counter
DECLARE @counter INT = 0;

-- Insert 100 records of (1, 5)
WHILE @counter < 100
BEGIN
    INSERT INTO ExPlanOperator_JOIN (STD_ID, STD_AbsenceDays)
    VALUES (1, 5);
    SET @counter = @counter + 1;
END
GO

-- Reset counter
SET @counter = 0;

-- Insert 100 records of (10, 2)
WHILE @counter < 100
BEGIN
    INSERT INTO ExPlanOperator_JOIN (STD_ID, STD_AbsenceDays)
    VALUES (10, 2);
    SET @counter = @counter + 1;
END
GO

SELECT STD_Name
  ,STD_Grade
  ,STD_AbsenceDays
FROM ExPlanOperator_P3 P3
INNER JOIN ExPlanOperator_JOIN AB ON P3.ID = AB.STD_ID
```

- How it works: The engine fills the hash table with data from the small table _JOIN_ (Probe) then process the larger table (Build)
	- Build phase: Each row of the larger table is hashed using the join key and inserted into an in-memory hash table

![[Pasted image 20240802142044.png]]