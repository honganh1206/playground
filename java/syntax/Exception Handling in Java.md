---
id: Exception Handling in Java
aliases:
  - Exception Handling in Java
tags: []
---

# Exception Handling in Java

[[Checked exceptions]]

```java
import java.io.*;

public class ExceptionExample {
    // Method that declares a checked exception
    public void readFile(String path) throws IOException {
        // In Java, IOException is a checked exception that must be declared or caught
        FileReader reader = new FileReader(path);
        BufferedReader bufferedReader = new BufferedReader(reader);

        String line;
        while ((line = bufferedReader.readLine()) != null) {
            System.out.println(line);
        }

        bufferedReader.close();
    }

    // Method that handles a checked exception
    public void safeReadFile(String path) {
        try {
            readFile(path);
        } catch (IOException e) {
            // Must handle IOException
            System.err.println("Error reading file: " + e.getMessage());
        } finally {
            // Cleanup code
            System.out.println("File operation attempt completed");
        }

        // Unchecked exceptions (similar to C#) don't need to be declared or caught
        String str = null;
        try {
            // This throws NullPointerException (unchecked, similar to C# NullReferenceException)
            int length = str.length();
        } catch (NullPointerException e) {
            System.err.println("Null reference: " + e.getMessage());
        }
    }

    // Try-with-resources (similar to C# using statement)
    public String readFirstLine(String path) throws IOException {
        try (BufferedReader reader = new BufferedReader(new FileReader(path))) {
            return reader.readLine();
        }
    }

    // Java 7+ multi-catch
    public void multiCatchExample() {
        try {
            // Code that might throw different exceptions
            File file = new File("nonexistent.txt");
            FileInputStream stream = new FileInputStream(file);
        } catch (FileNotFoundException | SecurityException e) {
            // Handle multiple exception types in one catch block
            System.err.println("File error: " + e.getMessage());
        }
    }
}
```
