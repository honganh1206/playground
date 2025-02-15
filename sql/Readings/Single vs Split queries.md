## Performance issue with single queries

`JOIN` can cause performance issues if used improperly

### Cartesian explosion

```cs
var blogs = ctx.Blogs
	// same-level collection navigations
    .Include(b => b.Posts) 
    .Include(b => b.Contributors) 
    .ToList();
```

```sql
SELECT [b].[Id], [b].[Name], [p].[Id], [p].[BlogId], [p].[Title], [c].[Id], [c].[BlogId], [c].[FirstName], [c].[LastName]
FROM [Blogs] AS [b]
LEFT JOIN [Posts] AS [p] ON [b].[Id] = [p].[BlogId] -- Join each blog to its corresponding posts => Each post listed with a blog
LEFT JOIN [Contributors] AS [c] ON [b].[Id] = [c].[BlogId] -- Join each contributor to its blog => Each contributor listed with a blog
ORDER BY [b].[Id], [p].[Id]
```

- In this example, SQL does not know the relationship between `[Posts]` and `[Contributors]` => 6 rows listed if we have a blog with 3 posts and 2 contributors

> [!info] Cartesian explosion
> 
> Cartesian explosion can cause **huge amounts of data to unintentionally get transferred to the client**, especially as more sibling JOINs are added to the query.


> [!note]
> 
> Cartesian explosion does not occur when the two `JOIN`s aren't at the same level.


### Data duplication

```sql
-- Each row contains properties form both Blogs and Posts tables
SELECT [b].[Id], [b].[Name], [b].[HugeColumn], [p].[Id], [p].[BlogId], [p].[Title]
FROM [Blogs] AS [b]
LEFT JOIN [Posts] AS [p] ON [b].[Id] = [p].[BlogId]
ORDER BY [b].[Id]
```

Supposed that the `Blogs` table a has a `HugeColumn` with binary data/huge text => Duplicated data sent back to clients multiple times

=> Solution: Use `SELECT` with to explicitly choose which columns to project

```cs
var blogs = ctx.Blogs
    .Select(b => new
    {
        b.Id,
        b.Name,
        b.Posts
    })
    .ToList();
```

## Split queries

 Instead of JOINs, split queries generate an additional SQL query **for each included collection navigation.**

```cs
using (var context = new BloggingContext())
{
    var blogs = context.Blogs
        .Include(blog => blog.Posts)
        .AsSplitQuery()
        .ToList();
}
```

```sql
SELECT [b].[BlogId], [b].[OwnerId], [b].[Rating], [b].[Url]
FROM [Blogs] AS [b]
ORDER BY [b].[BlogId]

SELECT [p].[PostId], [p].[AuthorId], [p].[BlogId], [p].[Content], [p].[Rating], [p].[Title], [b].[BlogId]
FROM [Blogs] AS [b]
INNER JOIN [Posts] AS [p] ON [b].[BlogId] = [p].[BlogId]
ORDER BY [b].[BlogId]
```


> [!note]
> 
> One-to-one related entities (each record in table A is connected to ONE AND ONLY ONE record in table B) are always loaded via JOINs in the same query, as it has no performance impact.


## Enable split queries

```cs
// Globally
protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
{
    optionsBuilder
        .UseSqlServer(
            @"Server=(localdb)\mssqllocaldb;Database=EFQuerying;Trusted_Connection=True;ConnectRetryCount=0",
            o => o.UseQuerySplittingBehavior(QuerySplittingBehavior.SplitQuery));
}

// Per-query
using (var context = new SplitQueriesBloggingContext())
{
    var blogs = context.Blogs
        .Include(blog => blog.Posts)
        .AsSingleQuery()
        .ToList();
}
```

## Characteristics of split queries

- No data consistency guarantee for split queries
- Each query implies additional network roundtrip => Might degrade performance
- As some databases only allow single query to be active at a given point, results from earlier queries must be buffered in your app’s memory => increased memory requirements
- Many reference navigations will lead to many JOINs => Degrade performance