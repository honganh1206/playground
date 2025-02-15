# Designing databases

## Data modeling

- Steps to take:
	1. Understand business requirement
	2. Build a conceptual model
	3. Build a logical model
	4. Build a physical model - Exact data types + technology

## Conceptual model

- Represent the **entities** and their **relationships**
- Represented with Entity Relationship (ER) or Unified Modeling Language (UML) diagram

![[Pasted image 20240803112939.png|Only represent business entities and relationships]]

## 3 main relationships

![[Pasted image 20240803112359.png]]


## Logical model

![[Pasted image 20240803113052.png]]


> [!tip] Designing tables with MySQL client
> 
> Choose `CASCADE` for On Update and `NO ACTION` for On Delete

## Normalization

- The process of **reviewing our design**
- Have seven rules/normal forms, each rule/form can only be applied if the previous rule/form is applied

> [!note]
> 
> Most business data model only requires the three first normal forms


### 1st Normal Form

- Each cell should have a single value
- We CANNOT have repeated columns

### 2nd Normal Form

- Every table should describe one entity, and every column in that table should describe that entity

### 3rd Normal Form

- All attributes in a table are determined only by the candidate keys of that relation and not by any non-prime attributes


> [!tip] Tips on designing databases
> 
> - Sometimes, logical analysis should be prioritized over normalization rules, so **we should always build the logical model before the physical one**.
> - Simple designs are always preferred => Solve today’s problems, not tomorrow’s problems that might never happen.


## Character Sets & Collations

> [!tip]
> 
> Most of the times we can use the default charset, but sometimes we must consider changing the charset to fit the specific storage size.


```sql
SHOW CHARSET -- in MySQL only
```