---
tags:
  - "#study"
cssclasses:
  - center-images
---
#1: Do nothing. A connection failure during transaction commit is low.

#2: Rebuild application state by discarding + creating a new `DbContext`

#3: Add state verification with `IExecutionStrategy.ExecuteInTransaction`

```cs
using var db = new BloggingContext();
var strategy = db.Database.CreateExecutionStrategy();

var blogToAdd = new Blog { Url = "http://blogs.msdn.com/dotnet" };
db.Blogs.Add(blogToAdd);

strategy.ExecuteInTransaction(
    db,
    operation: context => { context.SaveChanges(acceptAllChangesOnSuccess: false); },
	// only invoked when a transient error occurs during the transaction
    verifySucceeded: context => context.Blogs.AsNoTracking().Any(b => b.BlogId == blogToAdd.BlogId));

db.ChangeTracker.AcceptAllChanges();
```

#4: Manually track the transaction by assigning an ID for each transaction

```cs
using var db = new BloggingContext();
var strategy = db.Database.CreateExecutionStrategy();

db.Blogs.Add(new Blog { Url = "http://blogs.msdn.com/dotnet" });

var transaction = new TransactionRow { Id = Guid.NewGuid() }; // Add a new table
db.Transactions.Add(transaction); // Insert a row to a table

strategy.ExecuteInTransaction(
    db,
    operation: context => { context.SaveChanges(acceptAllChangesOnSuccess: false); },
    verifySucceeded: context => context.Transactions.AsNoTracking().Any(t => t.Id == transaction.Id));

db.ChangeTracker.AcceptAllChanges();
db.Transactions.Remove(transaction); // Remove the row if the transaction succeeds
db.SaveChanges();
```