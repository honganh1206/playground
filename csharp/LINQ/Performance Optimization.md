
## Use the right data structure

- `List<T>` for **small to medium-sized collections** where frequent additions/removals are needed
- `HashSet<T>` to ensure **fast lookup** and **unique elements**
- `Dictionary<TKey, TValue>` to store **key-value pairs** and perform **fast lookups** by key.

## Prefer compile-time query execution

- LINQ queries can be executed either at run-time (using `IEnumerable<T>`) or compile-time (using `IQueryable<T>`)
	- Compile-time: Query expressions are converted to method calls during compilation with type-safety in mind => Errors are caught earlier
	- Run-time: Queries are presented as expression trees then translated to SQL queries => More flexibility

```cs
// Using IEnumerable<T> (runtime query execution)
IEnumerable<Product> products = GetProducts();
var expensiveProducts = products.Where(p => p.Price > 1000);

// Using IQueryable<T> (compile-time query execution)
IQueryable<Product> products = GetProductsAsQueryable();
var expensiveProducts = products.Where(p => p.Price > 1000);
```

## Use the `Any` and `All` Methods Wisely

- Use `Any()` instead of `Count()` when checking if a collection has **at least one element**.
- Use `All()` to check if **all elements in a collection meet a specific condition**, instead of using `Where()` and `Count()`.

```cs
// Less efficient
bool hasItems = myCollection.Count() > 0;

// More efficient
bool hasItems = myCollection.Any();

// Less efficient
bool allItemsValid = myCollection.Where(x => x.IsValid).Count() == myCollection.Count();

// More efficient
bool allItemsValid = myCollection.All(x => x.IsValid);
```

## Prefer Lazy Evaluation

LINQ supports deferred execution, which means that **the query is not executed until the results are actually needed**.


```cs
// Eager evaluation (less efficient)
var resultList = myCollection.Where(x => x.IsValid).ToList();

// Lazy evaluation (more efficient)
IEnumerable<MyClass> result = myCollection.Where(x => x.IsValid);
```

## Use `Select` and `Where` Judiciously

- Use `Select` to **project only the necessary fields**, instead of returning the entire object.
- Chain `Where` clauses to **filter out unnecessary data early** in the query.

```cs
// Less efficient
var results = myCollection.Where(x => x.IsValid).Select(x => x);

// More efficient
var results = myCollection.Where(x => x.IsValid).Select(x => new { x.Id, x.Name });

// Less efficient
var results = myCollection.Where(x => x.IsValid).Where(x => x.IsActive);

// More efficient
var results = myCollection.Where(x => x.IsValid && x.IsActive);
```

## Leverage Parallel LINQ (PLINQ)

Parallel LINQ (PLINQ) enables you to execute LINQ queries in parallel, potentially improving performance for large data sets and CPU-bound operations.

> [!warning]
> 
> Ensure your machine’s CPU support parallel execution


```cs
var results = myCollection.AsParallel().Where(x => x.IsValid).Select(x => x.Name);
```

## Use Index-Based `Where` Overload

- The index-based overload allows you to filter items based on their index in the collection

```cs
// Using the standard Where method (less efficient)
var results = myCollection.Where(x => x.IsValid);

// Using the index-based Where method (more efficient)
var results = myCollection.Where((x, index) => x.IsValid && index % 2 == 0);
```


## Avoid Multiple Enumerations

Consider materializing the results into a concrete collection like `List<T>`or `Array<T>`

```cs
// Multiple enumerations (less efficient)
IEnumerable<MyClass> results = myCollection.Where(x => x.IsValid);
int count = results.Count();
foreach (var item in results) { /* ... */ }

// Materializing the results (more efficient)
List<MyClass> results = myCollection.Where(x => x.IsValid).ToList();
int count = results.Count;
foreach (var item in results) { /* ... */ }
```

## Optimize LINQ to SQL queries

- Use `Select` to project only the necessary fields from the database.
- Filter data using `Where` before applying other operations, such as `GroupBy` or `OrderBy`.
- Use `Take` and `Skip` for pagination instead of loading all the data and then filtering in the application.
- Use `CompiledQuery.Compile` to cache and reuse frequently executed queries.

```cs
// Less efficient
var allData = context.Products.ToList();
var filteredData = allData.Where(x => x.Price > 1000).Select(x => new { x.Id, x.Name });

// More efficient
var filteredData = context.Products.Where(x => x.Price > 1000).Select(x => new { x.Id, x.Name }).ToList();
```


## Use Predicate Builders for Dynamic Queries

```cs
// Using PredicateBuilder (more efficient)
var predicate = PredicateBuilder.True<MyClass>();
predicate = predicate.And(x => x.IsValid);
predicate = predicate.And(x => x.IsActive);
var results = myCollection.AsQueryable().Where(predicate);
```