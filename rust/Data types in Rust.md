## Arrays

Arrays in Rust own its elements, fixed length, and the size is part of the type

```rust
let arr = [10, 20, 30];
// type: [i32; 3]
```

`Vec` is a contiguous growable/shrinkable array type. It is heap-located and size is not part of the type (unlike array literals)



```rust
// This will compile
fn open(filename: &str) -> MyResult<Box<dyn BufRead>> {

// This will NOT compile since there is no fixed size of BufRead
fn open(filename: &str) -> MyResult<dyn BufRead> {
```

`Slices` as `&[T]` are like vectors but cannot be resized after creation. Usually to borrow data and just a view 

`Range` represents an interval between two bounds like `0..5`. Think of it like a short form representation of slices? like `0..5 == 0, 1, 2, 3, 4, 5`


## Type aliases

```rust
// Specific type of Result
// either an Ok with success type () returning nothing if succeeds
// or value that implements std::error::Error trait (A smart pointer that can hold any size)
type TestResult = Result<(), Box<dyn std::error::Error>>;
```


## Iterable with `Iterator::enumerate`

Return a tuple containing the index position and value for each element in an *iterable*, which can *produce values until exhausted*

`Values` is an iterator for getting multiple values out of an argument

## Enums

Rust's sum type `enum` is a *tagged union*. A tag indicates which variant is active.

The Rust compiler prevents invalid access because *it always knows the active variant*, unlike in C where we have to manually keep track of the active variant.

```rust
// Only one variant can have a value at a time
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

`Option` is a enum type that is either `None` or `Some<T>` where `T` is a type parameter.

The `Option<T>` enum has two variants:

- `None`, to indicate failure or lack of value, and
- `Some(value)`, a tuple struct that wraps a `value` with type `T`.

## Pointers

`Box` is a smart pointer that stores the value in the heap and gives ownership via a pointer-like handle


`Cow` is clone-on-write smart pointer. It encloses and provide immutable access to borrowed data, while cloning the data lazily when mutation or ownership is required.