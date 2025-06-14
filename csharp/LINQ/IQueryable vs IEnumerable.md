
[Link to a Medium article](https://medium.com/@shahrukhkhan_7802/understanding-the-differences-between-ienumerable-and-iqueryable-in-c-9dbc9ac145f3)

![[Pasted image 20240618111004.png]]

# `IEnumerable`

- Obtain an enumerator which allows **read-only** iteration over a collection
- Used mostly for in-memory data collection
## Key characteristics

- Client-side execution: Data is loaded to memory from the data source before filtering/sorting/etc.
- In-memory handling: Ideal to work with small-to-medium data collection e.g., arrays, lists
- Deferred execution: The execution of the query (the actual data processing part) will not happen when the query is defined, but rather when the data is enumerated

```csharp
IEnumerable<int> numbers = new List<int> { 1, 2, 3, 4, 5 };
IEnumerable<int> evenNumbers = numbers.Where(n => n % 2 == 0);
// At this point, there is no SQL query generated
// The data is already in RAM
// We just want to manipulate the data locally

foreach (int number in evenNumbers)
{
    Console.WriteLine(number);
}
```


- LINQ to objects: Data processed in app memory

# `IQueryable`

- An extension from `IEnumerable` 
- Query against a specific data source where the type of data is known => Preferred choice for remote data sources

## Key characteristics

- Server side execution: Allow conversing from LINQ queries to SQL
- Optimized for large datasets: Minimize the amount of data to be transferred
- Dynamic query generation: Able to create queries specific to the data source
- LINQ to SQL/Entities: Queries are processed in the database rather than the app memory

```cs
IQueryable<User> users = dbContext.Users;
var activeUsers = users.Where(u => u.IsActive);
var list = activeUsers.ToList(); // SQL sent **now**.
```

# Comparing `IEnumerable` and `IQueryable`

| `IEnumerable`         | `IQueryable`                |
| --------------------- | --------------------------- |
| Client-side execution | Server-side execution       |
| In-memory handling    | Query from the known source |
| Deferred execution    | Dynamic query generations   |
| LINQ to objects       | LINQ to SQL/Entities        |

# When to use? 

- Use `IEnumerable`:
	- For in-memory data collections.
	- When working with small to medium datasets.
	- When the data source is not a database.

- Use `IQueryable`:
	- For large datasets or databases.
	- When querying data from remote sources like a web service or a database.
	- When performance optimization and efficient data handling are required.