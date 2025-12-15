## Traits

Think of traits in Rust like interfaces in Go, but trait implementations are **explicit** instead of implicit like in Go:

```rust
impl Trait for Type {
    fn method(&self) { … }
}
```


## Bounds

The type parameters often must use traits as *bounds* to stipulate what functionality a type implements.

```rust
// Define a function `printer` that takes a generic type `T` which
// must implement trait `Display`.
fn printer<T: Display>(t: T) {
    println!("{}", t);
}

// Bounding restricts the generic to types that confirm to the bounds
struct S<T: Display>(T);

// Error! `Vec<T>` does not implement `Display`. This
// specialization will fail.
let s = S(vec![1]);
```

