---
tags:
  - "#study"
  - "#review"
  - "#programming"
cssclasses:
  - center-images
url: https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/reference-types
---

While variables of reference types store **references** to their data (objects), variables of value types **directly contain the data** with *each instance has its own copy*.

Some examples of value types are `int`, `float`, `struct` and `enum`. Examples of reference types are `class`, `string` and `delegate`

```cs
struct Point
{
    public int X;
    public int Y;
}

Point p1 = new Point { X = 10, Y = 20 };
Point p2 = p1;  // p2 gets a copy of p1's values
p2.X = 30;      // Modifying p2 doesn't affect p1

```

In the context of reference types, *two variables can refer to the same object*, so operations performed on one variable can *affect* the object referred to in another variable.

```cs
class Person
{
    public string Name;
}

Person p1 = new Person { Name = "Alice" };
Person p2 = p1;  // p2 now references the same object as p1
p2.Name = "Bob"; // Changing p2 also changes p1's Name, because both reference the same object

```

| Feature                | Value Types                                                            | Reference Types                                                |
| ---------------------- | ---------------------------------------------------------------------- | -------------------------------------------------------------- |
| Storage Location       | Typically stored on the stack                                          | Stored on the heap                                             |
| Memory Management      | Each variable gets a copy of the value                                 | Variables point to the same object                             |
| Lifetime               | Short-lived (stack memory is released automatically when out of scope) | Longer-lived (heap memory is managed by garbage collector)     |
| Behavior on Assignment | Copies the value                                                       | Copies the reference (both variables point to the same object) |
| Examples               | `int`, `float`, `struct`, `enum`                                       | `class`, `string`, `array`, `object`, `delegate`               |