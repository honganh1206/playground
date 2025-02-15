
- The Result tab is what differs the Actual and Estimated execution plan
- Sometimes, the optimizer will suggest missing index to enhance performance

![[Pasted image 20240709173452.png]]


- Reading flow: Left to right + Top to bottom
- Arrows: Direction and amount of data passed

![[Pasted image 20240709173615.png|200]]


> [!tip]
> 
> If:
> - The execution plan shows thick arrows
> - The number of the rows that are passed through the arrows is large
> - The arrow at the beginning of the plan and the number of rows passed through the last arrow to the `SELECT` statement and returned by the query is small
> 
> => **A scan operation is performed incorrectly to a table** or **an index that should be fixed**.

- Each operator will have a percentage assigned to show **the cost of that operator relative to the overall query cost**

![[Pasted image 20240709174913.png|All estimated cost is relative to the Clustered Index Seek]]


> [!tip] Tips
> 
> - We usually concentrate on **the subtree cost** of the operator that represents the execution tree.
> - For index scanning, we should focus on the **Ordered Boolean** value.

