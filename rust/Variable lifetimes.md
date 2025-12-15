
The **lifetime** refers to *how long a value is valid for borrowing throughout a program*.

Consider the following example where we use `&str` instead of `String`


```rust
#[derive(Debug)]
pub struct Config {
	in_file: &str,
	out_file: Option<&str>,
	count: bool,
}

pub fn get_args() -> MyResult<Config> {
	// Created inside the function, carrying references
	// and when the function returns, matches is dropped
	// i.e., data is freed (like Dobby)
	// but now we have dangling pointers (pointing to invalid data)
	let matches = App::new("uniq")
	...
	// At this point Config is holding references to dead memory (dangling pointers) and not owning the data
	Ok(Config{...})
}
```

If we use `&str` (fat pointer) instead of dynamically-sized type `String`, the compiler would complain:

```css
error[E0106]: missing lifetime specifier
--> src/lib.rs:11:14
|
11 |
 in_file: &str,
|
 ^ expected named lifetime parameter
|
help: consider introducing a named lifetime parameter
|
10 | pub struct Config<'a> {
11 |
 in_file: &'a str,
```

The problem: We are taking references to values from `matches`, which goes out of scope at the end of the function and is then *dropped*.

If we return a `Config` that *stores references to a dropped value* would lead to **dangling pointers**.

## Static

A quick way to get through the borrowing-checking hassle. Do NOT exploit it.

```rust
struct Test {
    input: &'static str,
    out: &'static str,
    out_count: &'static str,
}
```

`'static` means the data pointed to by the reference *lives for the remaining lifetime  of the running program*.