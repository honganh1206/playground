---
tags:
  - "#study"
cssclasses:
  - center-images
---
Using `FirstOrDefault()` could be a bad habit, and there are other alternatives like `First()`, `SingleOrDefault()` and `Single()`

[[About FirstOrDefault(), First(), SingleOrDefault() and Single()]]

[[Why FirstOrDefault() is a code smell]]

[[Why is FirstOrDefault() overused]]


> [!tip]
> Try to search through your code base for the first 10 places that use `FirstOrDefault()`. Some of them might be better off with `SingleOrDefault()`.


