---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#csharp"
cssclasses:
  - center-images
---
We can chain additional jobs once a task finishes its execution with `ContinueWith()`


```cs
string stringResult = await httpClient.GetStringAsync("https://ndepend.com/data")
					  .ContinueWith(task => CalculateComplexOutput(task.Result));

```

We use `Task.WhenAll()` if all tasks complete their executions (Task results are independent) and `Task.WhenAny()` to obtain the result of any task that completes first.


```cs
static async Task<int> GetNumberAfterDelay(int number, int delay)
{
    await Task.Delay(delay);
    return number;
}

static async Task ExampleWhenAll()
{
    // Define three asynchronous tasks
    Task<int> task1 = Task.Run(() => GetNumberAfterDelay(1, 1000));  // 1-second delay
    Task<int> task2 = Task.Run(() => GetNumberAfterDelay(2, 2000));  // 2-second delay
    Task<int> task3 = Task.Run(() => GetNumberAfterDelay(3, 1500));  // 1.5-second delay

    // Wait for all tasks to complete
    int[] results = await Task.WhenAll(task1, task2, task3);

    Console.WriteLine("All tasks completed.");
    foreach (int result in results)
    {
        Console.WriteLine($"Task result: {result}");
    }
}

static async Task ExampleWhenAny()
{
    // Define three asynchronous tasks
    Task<int> task1 = Task.Run(() => GetNumberAfterDelay(1, 3000));  // 3-second delay
    Task<int> task2 = Task.Run(() => GetNumberAfterDelay(2, 1000));  // 1-second delay
    Task<int> task3 = Task.Run(() => GetNumberAfterDelay(3, 2000));  // 2-second delay

    // Wait for the first task to complete
    Task<int> firstCompletedTask = await Task.WhenAny(task1, task2, task3);

    int result = await firstCompletedTask;  // Await the completed task to get its result
    Console.WriteLine($"First task completed with result: {result}");
}

await ExampleWhenAll(); // Exit when every task completes
await ExampleWhenAny(); // Exit when the 1-second delay task completes

```
