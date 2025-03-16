---
id: Java Collection Framework
aliases:
  - Java Collection Framework
tags: []
---

# Java Collection Framework

```java
import java.util.*;
import java.util.stream.*;

public class CollectionsExample {
    public void demonstrateCollections() {
        // List example (similar to C# List<T>)
        List<String> names = new ArrayList<>();
        names.add("Alice");
        names.add("Bob");
        names.add("Charlie");

        // Map example (similar to C# Dictionary<K,V>)
        Map<Integer, String> employees = new HashMap<>();
        employees.put(1001, "John Doe");
        employees.put(1002, "Jane Smith");

        // Set example (similar to C# HashSet<T>)
        Set<String> uniqueNames = new HashSet<>();
        uniqueNames.add("Dave");
        uniqueNames.add("Eve");

        // Queue example (similar to C# Queue<T>)
        Queue<String> waitingList = new LinkedList<>();
        waitingList.offer("Customer 1");
        waitingList.offer("Customer 2");
        String next = waitingList.poll(); // Retrieves and removes

        // Iteration using enhanced for loop
        for (String name : names) {
            System.out.println(name);
        }

        // Using Iterator (more explicit than C# foreach)
        Iterator<String> iterator = names.iterator();
        while (iterator.hasNext()) {
            String name = iterator.next();
            System.out.println(name);
        }
    }
}
```
