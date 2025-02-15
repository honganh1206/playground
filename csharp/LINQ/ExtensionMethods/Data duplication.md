---
tags:
  - "#study"
cssclasses:
  - center-images
---
If we don’t have constraints at the storage level, we will have duplicates and `SingleOrDefault()` will throw an exception => In this case we have to use `FirstOrDefault()`


> [!tip]
> - Try to introduce strong consistency guarantees at the database level.
> - Log a warning message if we encounter more than a single record when there should be exactly one record.




