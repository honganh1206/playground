---
tags:
  - "#study"
cssclasses:
  - center-images
---
# [Connection Resiliency]([Connection Resiliency - EF Core | Microsoft Learn](https://learn.microsoft.com/en-us/ef/core/miscellaneous/connection-resiliency))

The ability to **automatically retry failed database commands**. This can be used with **any** database by providing an “execution strategy” which helps us detect failures and retry commands.

```cs
// Enable execution strategy
public void ConfigureServices(IServiceCollection services)
{
    services.AddDbContext<PicnicContext>(
        options => options.UseSqlServer(
            "<connection string>",
            providerOptions => providerOptions.EnableRetryOnFailure()));
}

// Custom execution strategy
protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
{
    optionsBuilder
        .UseMyProvider(
            "<connection string>",
            options => options.ExecutionStrategy(...));
}
```

[[Execution strategies and transactions]]

[[Transaction commit failure and the idempotency issue]]