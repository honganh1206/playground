---
tags:
  - "#study"
  - "#review"
  - "#sql"
  - "#programming"
cssclasses:
  - center-images
---
Simply an index with **multiple key columns**, and can either be clustered or non-clustered. When all/most of all keys from left to right as specified in the index e.g., *idx_Key1_Key2* are in `WHERE/JOIN`, the index only *reads the rows needed by the query and avoids other*.

We can use composite index to **guarantee uniqueness (two rows not sharing combination) over multiple columns** thus ensuring data integrity.

