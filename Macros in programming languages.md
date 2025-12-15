A way of writing code that writes other code.

## The differences between macros and functions

A function signature must declare the number and type of parameters the function has, but macros can *take a variable number of parameters* like `println!("hello")` or `println!("hello {}", name)`.

Macros are expanded before the compiler interprets the meaning of the code, so a macro can *implement a trait on a given type*. Functions cannot do that, since they are called during runtime, while a trait needs to implemented at compile time.

However, macros are more complex, since we are using the current language to write mini versions of it.

Also, we need to define macros or bring them into scope *before* we call them in a file, while we can define functions and call them anywhere.

```rust

// Defined before being used, as opposed to functions
macro_rules! foo {
    () => {
        println!("Hello from macro!");
    };
}

foo!();   // now it works

```