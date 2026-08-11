Some programming languages struggle with memory-related issues like:

- Using memory after it was freed
- Two parts of a program changing the same data or one part reading, another part modifying (Race condition)

Rust tries to prevent these by *enforcing rules at compile time* instead of runtime.

Rust's **borrow-checker** works as a *strict librarian for memory*.

## Ownership

In C/C++, memory allocation works like this:

![[image.png]]

A `std::string` owns its buffer, when the program destroys the string, the string's de-structor frees the buffer.

While other code parts can create temporary pointers to an owned memory, those code parts are responsible for removing the pointers when the memory owner destroys the owned project.

In Rust, **every value has a single owner that determines its lifetime**. When the owner is freed/dropped, the owned value is dropped too.

Data structure ownership in Rust like arrays has a tree-like structure. We drop a value to remove the value from the ownership tree.

```rust
// s owns the string
// when s goes out of scope, the memory is freed automatically
let s1 = String::from("hi");
let s2 = s1;
// s1 is now invalid

```

We can use **reference counters** like `Rc` and `Arc` to allow values to have multiple owners (under some restrictions of course)
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

## Moves

In Rust we don't copy values when passing it to a function for example, **we move it**. The destination controls the value's lifetime.

How Python represents a list in memory:

![[image-1.png|647x466]]

Reference count (Specifically in Python) == Number of things pointing to the list.

Assignment to the other variables:

![[image-2.png]]

How C++ represents a vector of strings:

![[image-3.png]]

In C++, assigning a vector produces a copy of the original vector. When the variables go out of scope, the program can automatically clean up the allocations.

![[image-4.png]]


# `Copy` types

When we assign a value of a `Copy` type, we copy the value rather than moving it.

For simpler types like integers or characters, no need to move values around and we can just copy the values

![[image-5.png]]

Moving the value leaves the source of the value uninitialized.

Only simple bit-for-bit types can be copied like bool, float, integer, etc. This works with struct types wrapping simple types like `i32` and simple

> Copy types are VERY limited in which types they can contain, while non-Copy types can use heap allocation.

Copy leaves the source initialized.

## Shared ownership

The value should simply live until everyone is done using it.

We use `Rc` (reference counter, non-thread safe) and `Arc` (Atomic reference count, safe to share between threads) 