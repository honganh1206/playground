---
tags:
  - "#study"
  - "#review"
cssclasses:
  - center-images
---

```cs
static async Task<int> MethodAAsync()
{
    for (int i = 0; i < 5; i++)
    {
        ConsoleWriteLine($" A{i}");
        await Task.Delay(100);
        ConsoleWriteLine($" A throws exception");
        throw new ApplicationException("Boum");
    }
    int result = 123;
    ConsoleWriteLine($" A returns result {result}");
    return result;
}

...
try
{
    await taskA; // If this were taskA.Wait(), the execption will NOT be handled
    ConsoleWriteLine($"The result of taskA is {taskA.Result}");
}
catch (ApplicationException ex)
{
    ConsoleWriteLine($"{ex.GetType().ToString()} Msg:{ex.Message}");
}

Console.ReadKey();

```

The `await` keyword allows *preserving exception flow*, meaning any exception occurring in `taskA` are *directly propagated* to the `catch` block. This is why the application will throw the original `ApplicationException` during runtime. 

On the other hand, using `taskA.Wait()`, as a synchronous blocking call, will wrap the `ApplicationException` inside another exception like `AggregateException`