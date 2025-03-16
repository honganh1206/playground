---
id: Abstract, Virtual, and Override in Java
aliases:
  - Abstract, Virtual, and Override in Java
tags: []
---

# Abstract, Virtual, and Override in Java

Java uses the same abstract keyword with identical meaning and functionality as in C#:

```java
abstract class Shape {
    public abstract double calculateArea(); // Abstract method

    public void display() { // Regular method
        System.out.println("Area: " + calculateArea());
    }
}

```

Java does not have a virtual keyword because **methods in Java are virtual by default**.

Any non-private, non-final, non-static method can be overridden in subclasses.

```java
class Animal {
    // This method is implicitly virtual
    public void makeSound() {
        System.out.println("Some generic sound");
    }
}
```

Java uses the `@Override` annotation rather than the `override` keyword. It's optional but recommended for clarity and compiler checking:

```java
class Dog extends Animal {
    @Override
    public void makeSound() {
        System.out.println("Woof!");
    }
}
```

Java uses the `final` keyword to prevent method overriding, which is somewhat the opposite of C#'s virtual:

```java
class Animal {
    public final void breathe() {
        // This method cannot be overridden in subclasses
    }
}
```
