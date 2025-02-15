---
tags:
  - "#consume"
  - "#dotnet"
cssclasses:
  - center-images
---
[Conditional compilation symbols in .NET](https://bartwullems.blogspot.com/2024/09/conditional-compilation-symbols-in-c.html?utm_source=newsletter.csharpdigest.net&utm_medium=newsletter&utm_campaign=implementing-blocked-floyd-warshall-algorithm)

- **Preprocessor directives** allowing the compiler to include/omit portions of code based on certain conditions
- Examples: `#define`, `#undef` to define/undefine symbols within a file, or conditionals like `#if`, `#elif`, `#else` and `#endif`


```cs
#define PLATTELANDSLOKET

#if PLATTELANDSLOKET
    public Guid Id { get; init; }
#else
    public int Id { get; init; }
#endif
```

- We can define symbols in several places like in code/project properties/through VS UI/command line
- Best practices:
	- Minimize the use
	- Use at a high level rather than for individual statements or lines
	- Document the use