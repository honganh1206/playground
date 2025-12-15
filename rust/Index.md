Rust guarantees safety via **borrow checker** that tracks *which part of a program has safe access to different parts of memory*.

Rust is statically typed and uses structs to represent complex data types.

Variables in Rust are *immutable* by default

Rust uses *enumerated* and *sum* types that allow a function to return a `Result` that can be either an `Ok` or `Err`

```css
├── Cargo.lock // Exact version of dependencies
├── Cargo.toml
├── src
│   └── main.rs
├── target // Build artifacts
│   ├── CACHEDIR.TAG 
│   ├── debug
|	└── tmp
└── tests
	└── cli.rs
``` 

Start a new Rust project with `cargo new project-name`

Rust libraries are called *crates* using semantic version numbers of `major.minor.patch`. A major change indicates a breaking change in the API.

> [!tip]
> Make a `tests` directory alongside the `src` directory as a convention when testing.


Rust program will exit with the value zero by default.

`::` is the path separator in Rust, similar to `->` in C or `.` in other languages.

[[Raising exceptions in Rust]]

[[Testing in Rust]]

[[Handling dependencies in Rust]]

[[Data types in Rust]]

[[Common practices in Rust]]

[[Macros in Rust]]

[[Notations in Rust]]

[[Scoping rules]]

[[Functions in Rust]]

[[Generics in Rust]]