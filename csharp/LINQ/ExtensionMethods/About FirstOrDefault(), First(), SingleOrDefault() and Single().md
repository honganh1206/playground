---
tags:
  - "#study"
cssclasses:
  - center-images
---
```cs
// Signature
public static TSource FirstOrDefault<TSource>(this IEnumerable<TSource> source, Func<TSource, bool> predicate)

public class Person
{
    public int Id { get; set; }
    public string Name { get; set; }
}

private static IEnumerable<Person> People => new[]
{
    new Person{Id = 1, Name = "John"},
    new Person{Id = 2, Name = "Alice"},
    new Person{Id = 3, Name = "John"}
};

People.FirstOrDefault(p => p.Name == "John"); // John with Id 1
People.FirstOrDefault(p => p.Name == "Bob"); // null

```

- `FirstOrDefault()` gets to the first `John` without traversing the whole collection => A performance bonus

```cs
People.First(p => p.Name == "John"); // John with Id 1
People.First(p => p.Name == "Bob"); // throws System.InvalidOperationException

People.SingleOrDefault(p => p.Name == "Alice"); // Alice with Id 2
People.SingleOrDefault(p => p.Name == "Bob"); // null
People.SingleOrDefault(p => p.Name == "John"); // throws System.InvalidOperationException


People.Single(p => p.Name == "Alice"); // Alice with Id 2
People.Single(p => p.Name == "George"); // throws System.InvalidOperationException
People.Single(p => p.Name == "John"); // throws System.InvalidOperationException

```

- `First()` will throw an exception if there is no element that satisfies the conditions. For `Single()` and `SingleOrDefault()`, the exception will be thrown if **there is no element** or **more than 1 element satisfy the conditions**