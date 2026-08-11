**TLDR**: Macros (particularly in C) are to generate inline code (code that exists within the caller), so after compilation we don't have to move to where the function exists.

A way of writing code that writes other code. Macros manipulate code at compile time.

If a function can solve the problem, we should use the function. Macros are for where functions fundamentally cannot.

Supposed every function needs logging:

```c#
Console.WriteLine("Starting");
DoWork();
Console.WriteLine("Finished");

// Imagine using a function
LogExecution(DoWork());
```

In essence that is:

```c#
var result = DoWork();
LogExecution(result);
```

The `LogExecution()` function never gets the code `DoWork()` but just the result of it

```text
execute DoWork()
log start
log end
```

A function cannot insert statements around another piece of code

## Does the `Action` type in C# resolve this?

Conceptually yes, we use it like this:

```c#
void ExecuteWithLogging(Action action)
{
    Console.WriteLine("Starting");

    action();

    Console.WriteLine("Finished");
}
```

Here we treat behaviors as data, but here the passed in functions are compiled into objects (delegates), something that can be invoked later (callbacks).

C# `Action` can only be invoked during runtime, while with macros the compiler can see the code before compilation

```
Read source
      │
      ▼
Run macro
      │
      ▼
Generate CustomerRepository class
      │
      ▼
Compile everything
```

Macros are code-generation instead, and the goal is for the compiler to *see the generated code before it can compile the program*