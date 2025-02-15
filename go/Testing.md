---
id: Testing
aliases: []
tags: []
---

## Questions

1. Why do we use pointer type for `t *testing.T`?

- TLDR:

The pointer allows the test framework to:

1. Track test status across multiple function calls
2. Accumulate error messages
3. Control test execution flow
4. Maintain test logging and reporting

- **Mutation and State tracking**

  - `testing.T` struct needs to _maintain the state throughout the test_ e.g., pass/fail status, error messages, etc.
  - Methods on `tesing.T`
  - Using a pointer ensures all test methods work with the same instance and can modify its state

- **Efficiency:**

  - Passing by pointer is more efficient than passing by value, especially since `testing.T` might be a larger struct
  - Avoids unnecessary copying of the entire testing struct when calling methods

- **Standard Go Convention:**

```go
func TestSomething(t *testing.T) {
    // Test code here
    if something != expected {
        t.Errorf("got %v, want %v", something, expected)
    }
}

func TestExample(t *testing.T) {
    // If t wasn't a pointer, this would modify a copy
    // and the test framework wouldn't know about failures
    if someCondition {
        t.Error("Test failed")
    }
}
```

- **Method Requirements:**

```go
// These methods need to modify the test state
t.Error()   // Marks test as failed
t.Fatal()   // Marks test as failed and stops execution
t.Log()     // Logs test output
```

- **Interface Implementation:**
  - The testing package's internal implementation relies on pointer receivers
  - Methods are defined on `*testing.T`, not on `testing.T`

Without a pointer:

```go
// Wrong - modifications wouldn't persist
func TestWrong(t testing.T) {
    t.Error() // This would modify a copy, not the original
}
```

With a pointer (correct):

```go
// Correct - modifications affect the actual test instance
func TestCorrect(t *testing.T) {
    t.Error() // This properly records the test failure
}
```
