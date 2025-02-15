---
tags:
  - "#study"
cssclasses:
  - center-images
---
An execution strategy needs to *play back each operation in a retry block that fails*. That is, *every operation performed via EF Core is retriable*: Each query and call to `SaveChanges()` will be retried as a unit in case a transient occurs

We use `BeginTransaction()` to *treat each group of operations as a unit* and **everything** inside the transaction would need to be re-played should a failure occur (notified by the exception `InvalidOperationException`).

[[Invoking the execution strategy]]