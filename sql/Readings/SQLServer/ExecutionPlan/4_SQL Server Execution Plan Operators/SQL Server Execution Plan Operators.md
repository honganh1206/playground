# SQL Server Execution Plan Operators


```sql
-- Script for testing without indexing
CREATE TABLE ExPlanOperator
( ID INT IDENTITY (1,1),
  First_Name VARCHAR(50),
  Last_name VARCHAR(50),
  Address VARCHAR(MAX)
 )
 GO
 INSERT INTO ExPlanOperator VALUES ('AA','BB','CC')
 GO 1000
  INSERT INTO ExPlanOperator VALUES ('DD','EE','FF')
 GO 1000
```

- Without the clustered indexing, the engine will scan **row by row**

![[Pasted image 20240801111837.png|Even with a WHERE clause, the engine will still scan the entire table]]

## [[Clustered Index seek operators]]

## [[Non-clustered index seek operator]]

## [[Row-Identifier (RID) Lookup Operator]]

## [[Key Lookup Operator]]

## [[Sort Operator]]

## [[Aggregate Operator]]

## [[Scalar Operator]]

## [[Concatenation Operator]]

## [[Assert Operator]]

## [[Hash Match Operator]]

## [[Hash Match Aggregate Operator]]

## [[Merge Join Operator]]

## [[Nested Loops Join Operator]]



