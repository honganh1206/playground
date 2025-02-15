---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#csharp"
cssclasses:
  - center-images
---
- A special **construct** that allows you to *iterate/loop* through a collection of items (one at a time) without **having to load the entire collection into memory**
- Benefit: **Lazy generate and return elements** in a sequence

## `yield` keyword

- Used in the context of **iterators** to create an iterator block => Turn the method into an iterator + the method can use a `foreach` loop to iterate over *a sequence of values*

### Use cases

1. Create iterator method

```csharp
public IEnumerable<int> CountUpToTen()
{
	for (int i = 1; i <= 10; i++)
	{
		yield return 1;
	}
}
```

2. Lazy evaluation

- Key point: **Values are not generated all at once** => Each value is generated one at a time as you *iterate over the collection*
- Benefit: Save memory + Improve performance

3. Used in `foreach` loop

```csharp
foreach (int number in CountUpToTen())
{
    Console.WriteLine(number);
}
```