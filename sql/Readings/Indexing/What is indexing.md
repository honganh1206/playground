---
tags:
  - "#study"
  - "#review"
  - "#sql"
  - "#programming"
cssclasses:
  - center-images
---

- Index is a structure that holds **the field the index is sorting** (keyword in a book index) and **a pointer from each record (page number in a book index) to their corresponding record in the original table** where the data is actually stored. 

![[Pasted image 20240708111511.png|Ordered table and indexes allow faster query with binary search]]

- Indexes allow us to create sorted lists *without having to create all new sorted tables or scanning the entire table*.

![[Pasted image 20240708131718.png|While the data stored in Table has an incrementing ID, data is stored in the Index alphabetically]]