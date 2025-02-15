---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#computer"
cssclasses:
  - center-images
---

```cs
static void ConsoleWriteLine(string str)
{
    int threadId = Thread.CurrentThread.ManagedThreadId;
    Console.ForegroundColor = threadId == 1 ? ConsoleColor.White : ConsoleColor.Cyan;
    Console.WriteLine(
       $"{str}{new string(' ', 26 - str.Length)}   Thread {threadId}");
}

// Simulate a CPU-bound task
static async Task<int> MethodAAsync()
{
    for (int i = 0; i < 5; i++)
    {
        ConsoleWriteLine($" A{i}");
		// The delay yields control back to the thread pool
		// Without the delay, this would be a near-synchronous execution
        await Task.Delay(100); 
    }
    int result = 123;
    ConsoleWriteLine($" A returns result {result}");
    return result;
}

ConsoleWriteLine($"Start Program");

Task<int> taskA = MethodAAsync(); // Represent the process of running MethodAAsync()

// Synchronous operation executed when task A is delayed
for (int i = 0; i < 5; i++)
{
    ConsoleWriteLine($" B{i}");
    Task.Delay(50).Wait();
}

ConsoleWriteLine("Wait for taskA termination");

// From this point onward, another thread besides #1 is chosen to execute the remaining code
await taskA;

ConsoleWriteLine($"The result of taskA is {taskA.Result}");
Console.ReadKey();

```

![[Impact of Async-Await on the workflow.png|The result]]

Here is how things work:

1. When the `await` keyword is met for the 1st time (i = 0), the remaining code inside `MethodAAsync()` is then paused and the synchronous code for task B is to run
2. `A0` is printed then the control is given to the thread pool so *task B can run synchronously* while task A continues its execution with some background threads
3. As task B waits for only 50 mil secs while task A waits for 100 mil secs, task A will print one string when task B is done printing 2 strings. 
4. When task B is done executing, the remaining code of task A is executed
5. The code after `await taskA` is executed on a *random pool thread*

## [[Role of the async keyword]]

## [[Explaining the Async-Await workflow in detail]]

## [[Exception Handling in a CSharp Asynchronous Workflow]]



