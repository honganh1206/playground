---
tags:
  - "#study"
cssclasses:
  - center-images
---
```cs
var person = People.FirstOrDefault(p => p.Name == "John");
var name = person.Name;

// Might very well throw an NullReferenceException
// Use First() or Single() if we are sure that an user named "John" exists in the db
```