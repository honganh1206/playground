---
id: Abstract
aliases:
  - Abstract
tags: []
---

# Abstract

An abstract class is a class that **cannot be instantiated directly** and is **meant to be inherited by other classes**.

Abstract classes can contain both regular methods and abstract methods (methods without implementation).

```cs
abstract class Shape
{
    public abstract double CalculateArea(); // Abstract method - no implementation

    public void Display()  // Regular method with implementation
    {
        Console.WriteLine($"Area: {CalculateArea()}");
    }
}
```

Abstract and [virtual](./Virtual and Override.md) methods exist to support inheritance and polymorphism, fundamental aspects of object-oriented programming.
