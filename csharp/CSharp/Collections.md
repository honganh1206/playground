---
id: Collections
aliases:
  - Collections
tags: []
---

# Collections

Collections in C# are classes that store and manage groups of objects. The most common collections include:

- Lists: Ordered, resizable collections
- Dictionaries: Key-value pairs
- Sets: Unique elements
- Queues/Stacks: FIFO/LIFO structures

Collections often use these concepts together:

```cs
// Abstract collection base class
public abstract class CollectionBase<T>
{
    protected List<T> items = new List<T>();

    // Virtual method that can be overridden
    public virtual void Add(T item)
    {
        items.Add(item);
    }

    // Abstract method that must be implemented
    public abstract T GetItem(int index);
}

// Concrete implementation
public class CustomCollection<T> : CollectionBase<T>
{
    // Override the virtual method
    public override void Add(T item)
    {
        // Custom logic before adding
        Console.WriteLine($"Adding {item}");
        base.Add(item);
    }

    // Implement the abstract method
    public override T GetItem(int index)
    {
        if (index < 0 || index >= items.Count)
            throw new IndexOutOfRangeException();

        return items[index];
    }
}
```
