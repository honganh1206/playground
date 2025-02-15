---
tags:
  - "#study"
  - "#review"
  - "#dotnet"
  - "#programming"
cssclasses:
  - center-images
---
A positional record and a positional readonly record struct are **init-only**, that means *the data cannot be changed after object initialization*, but this can be overridden if needed.

This immutability helps when you need a *data-centric type* to be *thread-safe* or you need to ensure *the hash code created by the hash function remains the same* in the hash table.

However, the init-only properties **have shallow immutability**. You cannot change the *value of value-type properties* or the *reference of reference-type properties*, but *the data the reference-type property refers to can be changed!*

```cs
public record Person(string FirstName, string LastName, string[] PhoneNumbers);

public static void Main()
{
    Person person = new("Nancy", "Davolio", new string[1] { "555-1234" });
    Console.WriteLine(person.PhoneNumbers[0]); // output: 555-1234

    person.PhoneNumbers[0] = "555-6789";
    Console.WriteLine(person.PhoneNumbers[0]); // output: 555-6789
}
```

[[Non-destructive mutation in record]]

