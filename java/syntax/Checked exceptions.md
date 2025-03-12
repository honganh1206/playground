---
id: Checked exceptions
aliases:
  - Checked exceptions
tags: []
---

# Checked exceptions

Java has checked exceptions because its designers wanted to **enforce error handling at compile time**, ensuring that developers explicitly handle exceptional cases.

Checked exceptions must either be caught with a `try-catch` block or declared with `throws` in the method signature, making error handling explicit and predictable.

C#, on the other hand, does not have checked exceptions because **its designers believed they lead to unnecessary boilerplate code** and reduce flexibility.

Instead, C# relies on runtime exceptions and documentation to guide developers in handling errors appropriately. The philosophy in C# favors using unchecked exceptions (`System.Exception` and its subclasses) combined with structured error-handling patterns such as `try-catch`, logging, and exception propagation.

### Key Differences:

1. **Compile-Time vs. Runtime Enforcement**

   - **Java:** Forces handling of checked exceptions at compile time.
   - **C#:** Leaves exception handling to the developer's discretion at runtime.

2. **Boilerplate Code vs. Flexibility**

   - **Java:** Checked exceptions can result in verbose code (e.g., multiple `try-catch` blocks or `throws` declarations).
   - **C#:** Avoids excessive try-catch declarations, leading to cleaner method signatures.

3. **Philosophy of Error Handling**
   - **Java:** Treats exceptions as an integral part of method contracts.
   - **C#:** Assumes developers are responsible for knowing and handling exceptions as needed.

### Practical Impact:

- Java's checked exceptions improve robustness but can make APIs cumbersome.
- C#'s unchecked exceptions simplify method signatures but require discipline in handling errors properly.

### Examples

```java
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class JavaCheckedExceptionExample {
    // Method declares it throws IOException (checked exception)
    public static void readFile(String filePath) throws IOException {
        BufferedReader reader = new BufferedReader(new FileReader(filePath));
        System.out.println(reader.readLine());
        reader.close();
    }

    public static void main(String[] args) {
        try {
            readFile("test.txt");
        } catch (IOException e) { // Must handle the exception
            System.err.println("Error reading file: " + e.getMessage());
        }
    }
}
```

```cs
using System;
using System.IO;

class CSharpUncheckedExceptionExample
{
    // No 'throws' declaration is needed
    static void ReadFile(string filePath)
    {
        using (StreamReader reader = new StreamReader(filePath))
        {
            Console.WriteLine(reader.ReadLine());
        }
    }

    static void Main()
    {
        try
        {
            ReadFile("test.txt");
        }
        catch (IOException e) // Handling is optional
        {
            // No throwing exceptions here
            Console.WriteLine("Error reading file: " + e.Message);
        }
    }
}
```
