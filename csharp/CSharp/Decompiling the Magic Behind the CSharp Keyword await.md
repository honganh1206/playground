---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#csharp"
cssclasses:
  - center-images
---

```cs
      ...
ConsoleWriteLine("Wait for taskA termination");
await taskA;

Console.WriteLine(new System.Diagnostics.StackTrace());

ConsoleWriteLine($"The result of taskA is {taskA.Result}");
Console.ReadKey();

```

![[Pasted image 20241112134529.png]]

A lot of code is compiled by the C# compiler and then calls the [Task Parallel Library (TPL)](https://learn.microsoft.com/en-us/dotnet/standard/parallel-programming/task-parallel-library-tpl) to make this happen.

