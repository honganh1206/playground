---
tags:
  - "#study"
  - "#review"
  - "#programming"
  - "#csharp"
cssclasses:
  - center-images
---

| Characteristics | Class                                                                                                                            | Record                                                                                                                                                                                                                                              |
| --------------- | -------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Mutability      | Mutable by default, and need customized immutability mechanisms to be immutable                                                  | Immutable by default, but can be mutable with certain mechanism e.g., using `with` keyword                                                                                                                                                          |
| Value equality  | Two instances of classes can only be considered as equal **if they refer to the same object in memory**                          | Two instances of a record are considered equal **if their properties have the same values** and **of the same type**                                                                                                                                |
| Inheritance     |                                                                                                                                  |                                                                                                                                                                                                                                                     |
| Deconstruction  |                                                                                                                                  | Records provide deconstruction support, making it easy to extract values from an instance.                                                                                                                                                          |
| Use cases       | - Need for inheritance/identity/mutability<br>- Encapsulation<br>- Fine-grained control over how objects are created and managed | - Used as a lightweight, immutable data structure<br>- For scenarios where value equality/content equality makes more sense than reference equality (being the same instance).<br>- Concise syntax<br>- No need for complex inheritance hierarchies |

Note that classes can be immutable as well e.g., strings, but classes allow modifications as *they represent an OOP entity*, while records are just for *data storage* and have no functionality.