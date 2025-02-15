---
tags:
  - "#study"
cssclasses:
  - center-images
---
- Calculate aggregate expressions in the submitted query by grouping the values of an aggregated column (MIN, MAX, COUNT, AVG, SUM)
- The Stream Aggregate operator is fast due to the fact that it **requires the rows to be sorted based on the columns specified in the `GROUP BY`** clause before aggregating these values.

![[Pasted image 20240802134246.png|If the rows are not sorted, the engine will be forced to use the SORT operator]]