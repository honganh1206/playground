---
id: Java Streams
aliases:
  - Java Streams
tags: []
---

# Java Streams

```java
import java.util.*;
import java.util.stream.*;

public class StreamExample {
    public void demonstrateStreams() {
        List<Person> people = Arrays.asList(
            new Person("Alice", 30),
            new Person("Bob", 25),
            new Person("Charlie", 35),
            new Person("David", 28)
        );

        // Filtering (equivalent to C# Where)
        List<Person> adults = people.stream()
            .filter(person -> person.getAge() >= 30)
            .collect(Collectors.toList());

        // Mapping (equivalent to C# Select)
        List<String> names = people.stream()
            .map(Person::getName)  // Method reference syntax
            .collect(Collectors.toList());

        // Sorting (equivalent to C# OrderBy)
        List<Person> sortedByAge = people.stream()
            .sorted(Comparator.comparing(Person::getAge))
            .collect(Collectors.toList());

        // Reducing (equivalent to C# Aggregate)
        int totalAge = people.stream()
            .mapToInt(Person::getAge)
            .sum();

        // Grouping (equivalent to C# GroupBy)
        Map<Integer, List<Person>> peopleByAge = people.stream()
            .collect(Collectors.groupingBy(Person::getAge));

        // Parallel streams (similar to C# PLINQ)
        List<String> processedNames = people.parallelStream()
            .map(person -> processName(person.getName()))
            .collect(Collectors.toList());
    }

    private String processName(String name) {
        return name.toUpperCase();
    }

    // Person class for the example
    class Person {
        private String name;
        private int age;

        public Person(String name, int age) {
            this.name = name;
            this.age = age;
        }

        public String getName() { return name; }
        public int getAge() { return age; }
    }
}
```
