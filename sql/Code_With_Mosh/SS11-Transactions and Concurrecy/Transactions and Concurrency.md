

> [!info]
> 
> Transaction is **a group of SQL statements** that represent a single unit of work



> [!important]
> 
> MySQL wraps every of our SQL statements into transactions.


## Scenarios

### Lost updates

- Solution: Use locks

### Dirty reads

- Scenario: Transaction A rollbacks before transaction B completes 
- Solution: Read committed (Only read committed data)


### Non-repeating reads

- Scenario: At the time transaction B starts (like giving a discount), we should see the initial snapshot of read data from transaction A 

![[Pasted image 20240511121331.png]]


- Solution: Increase isolation level of transaction A (**Repeatable read**)

### Phantom read

- Scenario: Giving customers with more than 10 points a discount (read data from transaction A) but then after completing giving discount, one more customer becomes qualified for the discount as his/her point gets updated (transaction B).
- Solution: Increase isolation level with **Serializable** (transactions are aware of changes made to other transactions)



> [!warning] 
> 
> The higher the isolation level is, the more resources it takes to implement.


```sql
SHOW VARIABLES LIKE 'transaction_isolation';
SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
```


## Isolation levels 

![[Pasted image 20240511122508.png]]

### Read uncommitted

- Lowest isolation level => We may experience all concurrency problems

### Read committed

- Ensure a transaction can only read data that has been committed by other transactions

## Repeatable reads

If the transaction with the `SELECT` query completes has yet to be committed, the `SELECT` query will always return the pre-committed data even if the 2nd transaction `UPDATE` is committed

```sql
-- Session 1
-- Start a transaction
BEGIN TRANSACTION;

-- Read the balance for AccountId 1
SELECT Balance FROM Accounts WHERE AccountId = 1;
-- Result: 1000.00

-- The transaction is still open, and no changes have been made

-- Session 2
-- Start a transaction
BEGIN TRANSACTION;

-- Update the balance for AccountId 1
UPDATE Accounts
SET Balance = Balance - 500.00
WHERE AccountId = 1;

-- Commit the transaction
COMMIT TRANSACTION;

-- Back to session 1
-- Read the balance for AccountId 1 again
SELECT Balance FROM Accounts WHERE AccountId = 1;
-- Result: 1000.00 (The same as the first read)

-- Commit the transaction
COMMIT TRANSACTION;
```

### Serializable

If `UPDATE transaction is not yet committed, SELECT` transaction will wait for `UPDATE` to be committed


## Deadlocks

Both transactions wait for each other to be committed