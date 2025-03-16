---
id: Access modifiers
aliases:
  - Access modifiers
tags: []
---

# Access modifiers

```java
package com.example.interview;

// Public class - accessible from anywhere
public class AccessModifiersExample {
    // Public - accessible from any class
    public String publicField = "Public";

    // Protected - accessible within the same package and subclasses
    protected String protectedField = "Protected";

    // Package-private (no modifier) - accessible only within the same package
    String packagePrivateField = "Package Private"; // This doesn't exist in C#

    // Private - accessible only within the same class
    private String privateField = "Private";

    // Different access modifiers for methods
    public void publicMethod() { /* code */ }
    protected void protectedMethod() { /* code */ }
    void packagePrivateMethod() { /* code */ } // Package-private
    private void privateMethod() { /* code */ }
}
```
