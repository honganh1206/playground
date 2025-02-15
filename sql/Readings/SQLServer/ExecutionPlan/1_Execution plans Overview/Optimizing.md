---
tags:
  - "#study"
cssclasses:
  - center-images
---
- Input is the query tree into the Query Optimizer
- The optimizer **searches for the most efficient and cost-effective roadmap** to execute the submitted query + **ensures the optimization level settings is Full**
- The optimizer also **keeps the statistics of the database tables and indexes up-to-date** to create optimal execution plan

#### Plan Cache

- The optimizer will go through the cache before creating a new plan to save cost

> [!warning]
> 
> Having a large number of non-routine, spontaneous database queries (ad-hoc) will prevent the execution plans from being reused and require continuous plans generations.


#### Trivial plan

- Used to execute queries without aggregations or complex calculations