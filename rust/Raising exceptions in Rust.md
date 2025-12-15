Rust has recoverable errors, we use them with `Err`. For unrecovarable errors we use `panic`

```rust
// Result 
fn error_if_negative(value: i32) -> Result<(), &'static str> {
    if value < 0 {
        Err("Specified argument was out of the range of valid values. (Parameter 'value')")
    } else {
        Ok(())
    }
}
```

> Use `cargo run 1>out 2>err` to redirect output to out and err files.