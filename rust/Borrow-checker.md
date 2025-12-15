Some programming languages struggle with memory-related issues like:

- Using memory after it was freed
- Two parts of a program changing the same data or one part reading, another part modifying (Race condition)

Rust tries to prevent these by *enforcing rules at compile time* instead of runtime.

Rust's **borrow-checker** works as a *strict librarian for memory*.

## Ownership

```rust
// s owns the string
// when s goes out of scope, the memory is freed automatically
let s = String::from("hello");

let s1 = String::from("hi");
let s2 = s1;
// s1 is now invalid

```

## Borrowing

Use a value without taking ownership

```rust
let s = String::from("hello");
let r = &s; // borrow

```

The borrow-checker tracks:
- **Who owns the data**
- **Who is borrowing it**
- **For how long**

