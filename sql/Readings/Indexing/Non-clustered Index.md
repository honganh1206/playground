---
tags:
  - "#study"
  - "#review"
  - "#sql"
  - "#programming"
cssclasses:
  - center-images
sr-due: 2027-01-11
sr-interval: 473
sr-ease: 270
---
- Sorted references for **a specific field** from a main table
- Hold pointers back to the original entries of the table
- Created by DAs/devs after the table has been created

![[Pasted image 20240708134513.png|Increase the speed by creating more searchable columns]]


> [!warning]
> 
> Non-clustered indexes are NOT new tables.

- Non-clustered indexes point to **memory addresses** instead of storing data themselves, so they are slower than clustered index (not interacting with the real tables) but faster than a non-indexed column
- You can create as many non-clustered indexes as you can.