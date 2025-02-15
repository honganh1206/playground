---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#csharp"
cssclasses:
  - center-images
url: https://learn.microsoft.com/en-us/dotnet/csharp/advanced-topics/expression-trees/
---
- A data structure that **represents code in a tree-like format** where **each node is an expression** such as a method call or a binary operation 
- Can be examined/modified/executed at **runtime**
- Immutable
- Can be used as a **representation** of a lambda expression that can be manipulated programmatically
- Used in dynamic language runtime (DLR) to provide interoperability between dynamic languages and .NET

```cs
// Example of an expression tree representing a lambda expression
Expression<Func<int, bool>> exprTree = num => num < 5;
```

### Building an expression tree

#### Create nodes

```cs
// Leaf nodes
var one = Expression.Constant(1, typeof(int));
var two = Expression.Constant(2, typeof(int));
// Expression
var addition = Expression.Add(one, two);
// Lambda expression
var lambda = Expression.Lambda(addition)
```


#### Build a tree

```cs
var xParam = Expression.Parameter(typeof(double), "x")
var yParam = Expression.Parameter(typeof(double), "y")

var xSquared = Expression.Multiply(xParam, xParam)
var ySquared = Expression.Multiply(yParam, yParam)

var sum = Expression.Add(xSquared, ySquared)

// method call expression for the call to Math.Sqrt
var sqrtMethod = typeof(Math).GetMethod("Sqrt", new[] { typeof(double) }) ?? throw new InvalidOperationException("Math.Sqrt not found!");
var distance = Expression.Call(sqrtMethod, sum);

var distanceLambda = Expression.Lambda( distance, xParam, yParam);

```

## Components

- Constant expressions
- Parameter expressions
- Binary expressions -  Addition, multiplication
- Method call expression
- Lambda expression- Represent **the entire lambda expression** and can contain parameters and a body