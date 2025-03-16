---
id: Lambda Expressions and Functional Interfaces
aliases:
  - Lambda Expressions and Functional Interfaces
tags: []
---

# Lambda Expressions and Functional Interfaces

```java
import java.util._;
import java.util.function._;

public class LambdaExample {
public void demonstrateLambdas() {
// Functional interfaces in Java

        // Predicate - takes one argument, returns boolean (similar to C# Predicate<T>)
        Predicate<String> isLong = s -> s.length() > 5;
        boolean result = isLong.test("Hello, World!");  // true

        // Function - transforms a value (similar to C# Func<T, R>)
        Function<String, Integer> length = s -> s.length();
        int len = length.apply("Hello");  // 5

        // Consumer - accepts a value, returns nothing (similar to C# Action<T>)
        Consumer<String> printer = s -> System.out.println(s);
        printer.accept("Hello, Consumer!");

        // Supplier - provides a value (similar to C# Func<T>)
        Supplier<Double> random = () -> Math.random();
        double value = random.get();

        // BiFunction - takes two arguments, returns a value
        BiFunction<Integer, Integer, Integer> add = (a, b) -> a + b;
        int sum = add.apply(5, 3);  // 8

        // Method references
        List<String> names = Arrays.asList("Alice", "Bob", "Charlie");

        // Lambda form
        names.forEach(name -> System.out.println(name));

        // Method reference form (similar to C# delegate)
        names.forEach(System.out::println);
    }

    // Custom functional interface (similar to C# delegate or interface with single method)
    @FunctionalInterface  // Optional annotation ensures interface has exactly one abstract method
    interface Calculator {
        int calculate(int a, int b);

        // Can have default methods (Java 8+)
        default void describe() {
            System.out.println("A simple calculator interface");
        }
    }

    public void useCalculator() {
        // Lambda implementing the Calculator interface
        Calculator adder = (a, b) -> a + b;
        Calculator multiplier = (a, b) -> a * b;

        System.out.println(adder.calculate(5, 3));      // 8
        System.out.println(multiplier.calculate(5, 3)); // 15

        // Using default method
        adder.describe();
    }

}
```
