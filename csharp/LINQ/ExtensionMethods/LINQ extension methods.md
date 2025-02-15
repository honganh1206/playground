
# Improvements in .NET 8

## [OPENJSON](https://devblogs.microsoft.com/dotnet/announcing-ef8-preview-4/)

## [LINQ in general](https://devblogs.microsoft.com/dotnet/performance-improvements-in-net-8/#linq)

Key takeaways:

- Most systems nowadays support 128-bit vectors

## How things work

1. We run the exact same SQL to the database => EF sends the exact same SQL to the db
2. SQL Server caches SQL, performing expensive query planning only the first time a particular SQL is seen
## Issues

- Constantly varying SQLs: As SQL server/Npgsql can only cache a certain number of SQLs and need to remove old entries, using `Contains()` with a variable array causes **valuable cache entries to be taken at the db for SQLs that are most likely never be used**

## Solutions & Improvements

- `OPENJSON`: A “table-valued function” which accepts a JSON document, and returns a standard, relational rowset from its contents.

```sql
SELECT * FROM OPENJSON('["one", "two", "three"]'); -- Convert .NET array variable into a JSON array
```

=> Result:

|[key]|value|type|
|---|---|---|
|0|one|1|
|1|two|1|
|2|three|2|

> [!quote]
> 
> Importantly, when viewed on its own, **this new translation may actually run a bit _slower_ than the previous one** – SQL Server can sometimes execute the previous IN translation more efficiently than it can the new translation; when exactly this happens depends on the number of elements in the array.

- EF now supports querying into **any kind of primitive collection**, be it a column, a parameter or an inline collection.

```csharp
public class Blog
{
    public int Id { get; set; }
    // ...
    public string[] Tags { get; set; }
}

// Example 1
var blogs = await context.Blogs
    .Where(b => b.Tags.Contains("Tag1"))
    .ToArrayAsync();

// Example 2
var tags = new[] { "Tag1", "Tag2" };

//checks if the count of tags that intersect (common to both the blog's tags and the `tags` array) is greater than or equal to 2. 
// This means that a blog must have at least two tags from the `tags` array to be included in the result.
var blogs = await context.Blogs
    .Where(b => b.Tags.Intersect(tags).Count() >= 2)
    .ToArrayAsync();
```

```sql

-- Query for example 1
SELECT [b].[Id], [b].[Name], [b].[Tags]
FROM [Blogs] AS [b]
WHERE EXISTS (
    SELECT 1
    FROM OPENJSON([b].[Tags]) AS [t]
    WHERE [t].[value] = N'Tag1') -- Queryable primitive collections

-- Query for example 2
Executed DbCommand (48ms) [Parameters=[@__tags_0='["Tag1","Tag2"]' (Size = 4000)], CommandType='Text', CommandTimeout='30']

SELECT [b].[Id], [b].[Name], [b].[Tags]
FROM [Blogs] AS [b]
WHERE (
    SELECT COUNT(*)
    FROM (
        SELECT [t].[value]
        FROM OPENJSON([b].[Tags]) AS [t] -- column collection
        INTERSECT
        SELECT [t1].[value]
        FROM OPENJSON(@__tags_0) AS [t1] -- parameter collection
    ) AS [t0]) >= 2
```
# Methods
## `Include()`

- Specify the **navigation property** to be included in the query result => Load the specified navigation property along with the main entity
- Can be used multiple times
- If you used `Include()` for multiple navigation properties, they are loaded **in parallel**

```csharp
var result = dbContext.Orders
                .Include(o => o.Customer)
                .Include(o => o.OrderDetails)
                .ToList();

```


## `ThenInclude()`

- Specify **additional related data** to be included for a **previously included navigation property**
- Used after `Include()` to include **further levels of related entities**
- Can be used multiple times

```csharp
var result = dbContext.Orders
                .Include(o => o.Customer)
                    .ThenInclude(c => c.Address)
                .Include(o => o.OrderDetails)
                    .ThenInclude(od => od.Product)
                .ToList();

```


## [`Except()`](https://learn.microsoft.com/en-us/dotnet/api/system.linq.enumerable.except?view=net-8.0)

Produces the **set difference** of two sequences by using the default equality comparer to compare values.

```cs
double[] numbers1 = { 2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5 };
double[] numbers2 = { 2.2 };

IEnumerable<double> onlyInFirstSet = numbers1.Except(numbers2);

foreach (double number in onlyInFirstSet)
    Console.WriteLine(number);

/*
 This code produces the following output:

 2
 2.1
 2.3
 2.4
 2.5
*/
```

## [`Distict()`](https://learn.microsoft.com/en-us/dotnet/api/system.linq.enumerable.distinct?view=net-8.0)

Returns distinct elements from a sequence.

```cs
List<int> ages = new List<int> { 21, 46, 46, 55, 17, 21, 55, 55 };

IEnumerable<int> distinctAges = ages.Distinct();

Console.WriteLine("Distinct ages:");

foreach (int age in distinctAges)
{
    Console.WriteLine(age);
}

/*
 This code produces the following output:

 Distinct ages:
 21
 46
 55
 17
*/
```


## [`Intersect()`](https://learn.microsoft.com/en-us/dotnet/api/system.linq.enumerable.intersect?view=net-8.0)

Produces the set intersection of two sequences.

```cs
int[] id1 = { 44, 26, 92, 30, 71, 38 };
int[] id2 = { 39, 59, 83, 47, 26, 4, 30 };

IEnumerable<int> both = id1.Intersect(id2);

foreach (int id in both)
    Console.WriteLine(id);

/*
 This code produces the following output:

 26
 30
*/
```

## `Contains()`

### `Contains()` with inline collection

```csharp
var blogs = await context.Blogs
    .Where(b => new[] { "Blog1", "Blog2" }.Contains(b.Name)) // Names embedded as constants
    .ToArrayAsync();
```

```sql
SELECT [b].[Id], [b].[Name]
FROM [Blogs] AS [b]
WHERE [b].[Name] IN (N'Blog1', N'Blog2') -- Names embedded as constants
```

### `Contains()` with parameter collection

```csharp
var names = new[] { "Blog1", "Blog2" };

var blogs = await context.Blogs
    .Where(b => names.Contains(b.Name)) // Embed a variable within a query
    .ToArrayAsync();
```

```sql
SELECT [b].[Id], [b].[Name]
FROM [Blogs] AS [b]
WHERE [b].[Name] IN @names -- (INVALID) EF sends the variable as it is via a database parameter
WHERE [b].[Name] IN (N'Blog1', N'Blog2') -- The above C# code will be translated to this
```

> [!info]
> 
> Relational databases don’t really have the concept of a “list” or of a “collection”; they generally work with logically unordered, structured sets such as tables.


