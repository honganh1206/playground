---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#csharp"
cssclasses:
  - center-images
---

> [!summary] Key takeaways
> - `async/await` are used to *simplify the working* with CPU-bound and I/O-bound jobs
> - We denote async jobs with `Task` and `Task<T>` objects
> - When encountering the `await` keyword, the C# compiler creates a **state machine** to suspend execution → execute the async task → resume the execution once the async task is completed



We mark a method with the `async` keyword to denote that *the method executes non-blocking operations*. The keyword `await` in an async method causes the method to *return a task* - a representation of the remaining code to be executed until completion.

[[CPU-Bound]]

[[Input-Output Bound]]

[[POV from the caller method]]

When we use the keywords `async/await`, the C# compiler *converts the usage of the keyword `await` into a state machine*. This state machine will call the `Task<T>`-based APIs to:
- **Suspend** execution when encountering an `await` operation: By doing this, *the control is returned to the calling code*, and other code can be run while the asynchronous operation completes in the background
- **Execute** the CPU-bound (background thread) or I/O-bound (no thread occupation, return to original thread when completed) task 
- **Resume** to executing instructions after `await` when the task is completed

[[The Task Library from .NET base class]]

