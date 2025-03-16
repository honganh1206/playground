---
id: Generics in Java
aliases:
  - Generics in Java
tags: []
---

# Generics in Java

```java
// Generic class (similar to C# generic class)
public class Box<T> {
    private T content;

    public Box(T content) {
        this.content = content;
    }

    public T getContent() {
        return content;
    }

    public void setContent(T content) {
        this.content = content;
    }

    // Generic method
    public <U> void inspect(U item) {
        System.out.println("T: " + content.getClass().getName());
        System.out.println("U: " + item.getClass().getName());
    }
}

// Bounded type parameters
public class NumberBox<T extends Number> {
    private T number;

    public NumberBox(T number) {
        this.number = number;
    }

    public double square() {
        return number.doubleValue() * number.doubleValue();
    }
}

// Wildcard usage (? is similar to C# out/in keywords but with different syntax)
public class WildcardExample {
    // ? extends T (similar to C# out T)
    public double sumOfList(List<? extends Number> list) {
        double sum = 0;
        for (Number number : list) {
            sum += number.doubleValue();
        }
        return sum;
    }

    // ? super T (similar to C# in T)
    public void addNumbers(List<? super Integer> list) {
        list.add(1);
        list.add(2);
        list.add(3);
    }
}
```
