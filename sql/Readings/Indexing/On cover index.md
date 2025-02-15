---
tags:
  - "#study"
  - "#review"
  - "#sql"
  - "#programming"
cssclasses:
  - center-images
---
A cover index is when *all table columns needed by the query are present in the index*. Cover index *avoids additional reads* to retrieve values from actual tables, as *values from explicitly included columns are included in the index already*.

Columns available in non-clustered indexes:
1. Index key column(s)
2. Explicitly included column(s)
3. Clustered index key column(s)
4. Partitioning column(s) - implicitly included when table is partitioned

