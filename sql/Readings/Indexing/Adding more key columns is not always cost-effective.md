---
tags:
  - "#study"
  - "#review"
  - "#sql"
  - "#programming"
cssclasses:
  - center-images
---
Adding more key columns comes with the cost of *maintaining b-tree every time values in the `Quantity` changes*. Instead we could include the `Quantity` column

```sql
CREATE NONCLUSTERED INDEX idx_ProductInventory_Location
    ON Production.ProductInventory(LocationID) INCLUDE(Quantity);
```

