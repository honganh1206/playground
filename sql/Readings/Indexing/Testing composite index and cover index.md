---
tags:
  - "#study"
  - "#review"
  - "#sql"
  - "#programming"
cssclasses:
  - center-images
---
```sql
-- A table with a composite primary key index
CREATE TABLE [Production].[ProductInventory](
    [ProductID] [int] NOT NULL,
    [LocationID] [smallint] NOT NULL,
    [Shelf] [nvarchar](10) NOT NULL,
    [Bin] [tinyint] NOT NULL,
    [Quantity] [smallint] NOT NULL,
    [rowguid] [uniqueidentifier] ROWGUIDCOL  NOT NULL,
    [ModifiedDate] [datetime] NOT NULL,
    CONSTRAINT [PK_ProductInventory_ProductID_LocationID] PRIMARY KEY CLUSTERED 
    (
        [ProductID] ASC,
        [LocationID] ASC
    )
);
```

The index above only works *if the query specify AT LEAST the `ProductID`*

```sql
SELECT  
      ProductID
FROM Production.ProductInventory
WHERE 
    LocationID = @LocationID -- This will do a full scan, but new versions of MSSQL have now use clustered index
    AND Quantity = 0;
```

![[{B57A6855-C2C5-488A-9810-A2725808B2BD}.png|Version is SQL Server 15 - This is when we both add and not yet add non-clustered index]]


Adding `Quantity` as a key column helps when querying with `Quantity` column in the `WHERE/JOIN` condition

```sql
CREATE NONCLUSTERED INDEX idx_ProductInventory_Location_Quantity
    ON Production.ProductInventory(LocationID, Quantity);
```

![[{35E1E0CE-86A6-4720-9F80-619F51027A8F}.png|Better performance when adding another key column]]


[[Adding more key columns is not always cost-effective]]