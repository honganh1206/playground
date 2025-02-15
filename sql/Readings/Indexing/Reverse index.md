---
id: Reverse index
aliases: []
tags: []
---

# Reverse index

A type of index that stores a record of where search terms e.g., words and numbers. are located in a table

Example: A table of documents and we need to create a reverse index for those documents. The index needs to identify each document which contains a *particular word* using SQL  

```sql
CREATE TABLE docs01 (
    id INT IDENTITY(1,1) PRIMARY KEY,  -- SERIAL becomes INT IDENTITY
    doc TEXT
);

CREATE TABLE invert01 (
    keyword TEXT,
    doc_id INT FOREIGN KEY REFERENCES docs01(id) ON DELETE CASCADE -- INTEGER becomes INT, and add FOREIGN KEY constraint explicitly
);

INSERT INTO docs01 (doc) VALUES
('The building blocks of programs'),
('In the next few chapters we will learn more about the vocabulary'),
('sentence structure paragraph structure and story structure of Python'),
('We will learn about the powerful capabilities of Python and how to'),
('compose those capabilities together to create useful programs'),
('There are some lowlevel conceptual patterns that we use to construct'),
('programs These constructs are not just for Python programs they are'),
('part of every programming language from machine language up to the'),
('file or even some kind of sensor like a microphone or GPS In our'),
('initial programs our input will come from the user typing data on');

SELECT TOP 10 keyword, doc_id FROM invert01 ORDER BY keyword, doc_id;
```
We then split each character and index each of them with the document's id

```sql
INSERT INTO invert01 (keyword, doc_id)
-- Remove duplicates
SELECT DISTINCT LOWER(value) AS keyword, id AS doc_id
FROM docs01
CROSS APPLY STRING_SPLIT(CAST(doc AS NVARCHAR(MAX)), ' ');

ALTER TABLE invert01
ALTER COLUMN keyword NVARCHAR(MAX);  -- Change the data type
```

Then we can query the `invert01` table: `SELECT TOP 10 keyword, doc_id FROM invert01 ORDER BY keyword, doc_id;`
