---
tags:
  - "#consume"
  - "#sql"
cssclasses:
  - center-images
url: https://learn.microsoft.com/en-us/sql/relational-databases/indexes/create-indexes-with-included-columns?view=sql-server-2017
---
Including non-key columns allows you to *create non-clustered indexes that cover more queries*

An index with non-key columns can **significantly improve** query performance when *all columns in the query are included in the index as key/non-key column*

> An index containing all columns referenced by a query is referred to as **covering the query**

## Example

Supposed we have an `Orders` table with the following columns: `OrderID`, `CustomerID`, `OrderDate`, and `TotalAmount`. We can create an index on the `OrderID`: 

```sql
CREATE NONCLUSTERED INDEX IX_Orders ON Orders (OrderID) INCLUDE (CustomerID, TotalAmount);
```

So a query like `SELECT OrderID, CustomerID, TotalAmount FROM Orders WHERE OrderID = 123;` can now be **directly executed** from the index without having to access the real `Orders` table.

## Design recommendations

1. Keep the index key small and efficient. Make columns that cover the query non-key columns

```sql
CREATE NONCLUSTERED INDEX idx_Employees_Department
ON Employees (DepartmentID)
INCLUDE (FirstName, LastName, Salary);
```

2. Include non-key columns to avoid exceeding 32-key column limit and maximum index key size. *The Database Engine does not consider non-key columns when calculating key column number + index key size*
3. The order of non-key columns does NOT affect performance (But the order of key columns DOES!)
4. Avoid wide non-clustered indexes (indexes including many columns)