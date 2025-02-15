---
tags:
  - "#study"
  - "#review"
cssclasses:
  - center-images
url: https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/builtin-types/record
---
`record` aka `record class` is new reference type for encapsulating data, and `record` is a modifier

In a record, primary constructor parameters are called **positional parameters**, and the compiler creates **positional properties** that mirror the positional parameters.

But why? Conciseness I guess. We can reduce the boilerplate code by writing a single line record declaration.

```cs
public record Person(string FirstName, string LastName);
```


We can also have `record struct` value type:

```cs
public readonly record struct Point(double X, double Y, double Z);
```

Here is a way to initialize the property values when creating an instance with `record`:

```cs
public record Person(string FirstName, string LastName);

public static void Main()
{
    Person person = new("Nancy", "Davolio");
    Console.WriteLine(person);
    // output: Person { FirstName = Nancy, LastName = Davolio }
}
```


> [!important] 
> `record` and `readonly record struct` are **init-only** but **record struct** is read-write.

`record` does not require us to declare any positional properties, but we still can declare other fields and properties

```cs
public record Person(string FirstName, string LastName)
{
    public string[] PhoneNumbers { get; init; } = [];
};
```

[[Immutability of record]]

[[Value equality of record]]

