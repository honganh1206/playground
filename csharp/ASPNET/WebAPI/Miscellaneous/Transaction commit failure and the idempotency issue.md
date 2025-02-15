---
tags:
  - "#study"
cssclasses:
  - center-images
---
When there is a connection failure, *the current transaction is rolled back*. If so, *the execution strategy will retry the operation*, but this might lead to
- ⚠️ **An exception** *if the new database state is incompatible* 
- ⚠️ **Data corruption** *if the operation does not rely on a particular state*

[[Ways to resolve transaction failure and idempotency issues]]