
> [!tip] Metaphors for async programming
> Imagine threads as waiters. While the cooks are preparing fancy dishes (resource-heavy tasks), asynchronous programming allows the waiter/thread #1 to be freed from the current task and prepare for another task. When the resource-heavy task is completed, another waiter will be called in to serve. 


### A. The caller’s point of view

The caller of the method `MethodAAsync()` receives a **promise of the result** in the form of `Task<TResult>`, while the currently executing thread (thread #1) returns to the thread pool. The #1 thread then executes task B and then awaits the task A later when it needs the result. 

### B. The awaited asynchronous task

The keyword `await` is followed by a `Task`-typed object. This object could be either a simulated CPU-bound task `Task.Delay()` or already running in `Main()` as `taskA`

### C. The task returned by the async method is the remaining code once the awaited task terminates

The **catch**: The keyword `await` does NOT lead to wasted thread awaiting the task to end.

The remaining code of task A is nested with a `Task` object. When task B finishes, the infrastructure chooses another thread (thread #7 and thread #4) to execute the remaining code in the async method `MethodAAsync()`

There are **at least 2 tasks** involved in the `MethodAAsync()` method:

1. The CPU-bound task
2. The task returned by the `async` method it self representing the remaining code to be run

In each iteration, *a new task is created* to run the remaining code once the `await Task.Delay(100)` terminates. For this reason, the `taskA` returned by `MethodAAsync()` in `Task<int> taskA = MethodAAsync();` is actually **a chain of tasks** executed sequentially. Each loop is executed by a *randomly chosen thread* and such threads run the *sub-tasks* of `taskA`.