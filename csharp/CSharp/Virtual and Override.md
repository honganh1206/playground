---
id: Virtual and Override
aliases:
  - Virtual and Override
tags: []
---

# Virtual and Override

A virtual method has an implementation in the base class but can be overridden in derived classes. It allows for polymorphic behavior.

```cs
class Animal
{
    public virtual void MakeSound()
    {
        Console.WriteLine("Some generic sound");
    }
}
```

The override keyword is used in a derived class to provide a specific implementation for a method that is already defined in the base class as virtual or abstract.

```cs
class Dog : Animal
{
    public override void MakeSound()
    {
        Console.WriteLine("Woof!");
    }
}
```

Many collection classes (like `List<T>`, `Dictionary<K,V>`) use virtual methods internally to allow for customization in derived classes.
