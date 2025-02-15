# Execution plans

## Overview

- When you submit a Transact-SQL (T-SQL) query, you ask the SQL Server Engine (SSE) to **provide you with what you want without specifying how**.
- Before returning data to the end user, SSE will perform four internal query processing operations: Parsing/Algebrizing/Optimizing/Execution

![[Pasted image 20240708155126.png|400]]

> [!tip]
> 
> If the submitted query is not a Data Manipulation Language (DML) statement, such as `CREATE TABLE` or `ALTER TABLE`, there will be **no need to optimize** that query, as there is only one straight way for the SQL Server Engine to perform that action.

## [[Components]]

