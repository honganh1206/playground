---
id: Lambda functions and Delegate type
aliases: []
tags:
  -  #study
  -  #review
  -  #programming
  -  #csharp
cssclasses:
  - center-images
url: https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/operators/lambda-expressions
---

- A way to represent an **anonymous function** that can be used to **create delegates or expression tree types**
- Two forms:
  - **Expression lambda** - Expression as body `(input-parameter) => expression`
  - **Statement lambda** - A statement block as body `(input-parameter) => {<sequence-of-statements>}`
- `Func<T, object>` takes one parameter of type `T` and return a `object` type

```cs
// x as the input parameter
// x => x * x is a lambda funtion assigned to delegate of type Func<int, int>
Func<int, int> square = x => x * x;
Console.WriteLine(square(5));
Func<int, string, bool> isTrue = (n, s) => n == s; // Two input params and return type of bool
```

> [!tip]
>
> When writing lambdas, you often don't have to specify a type for the input parameters because **the compiler can infer the type** based on the lambda body, the parameter types, and other factors as described in the C# language specification.
