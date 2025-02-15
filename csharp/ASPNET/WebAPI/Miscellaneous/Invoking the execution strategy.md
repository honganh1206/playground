---
tags:
  - "#study"
cssclasses:
  - center-images
---
We *manually invoke the execution strategy* with a delegate that *represents everything that needs to be executed*. In case of failure, the execution strategy will invoke the delegate again.

```cs
using var db = new BloggingContext();
var strategy = db.Database.CreateExecutionStrategy();

strategy.Execute(
    () =>
    {
        using var context = new BloggingContext();
        using var transaction = context.Database.BeginTransaction();
        
		// If one of these fails, the delegate will be invoked again
        context.Blogs.Add(new Blog { Url = "http://blogs.msdn.com/dotnet" });
        context.SaveChanges();

        context.Blogs.Add(new Blog { Url = "http://blogs.msdn.com/visualstudio" });
        context.SaveChanges();

        transaction.Commit();
    });
```

[[This approach can be used with ambient transactions]]