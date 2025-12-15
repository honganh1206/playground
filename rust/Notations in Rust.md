The `?` notation in Rust is a concise way to handle `Result` and `Option` types, propagating errors or `None` values up the call stack

```rust
fn divide(numerator: f64, denominator: f64) -> Result<f64, String> {
    if denominator == 0.0 {
        Err("Cannot divide by zero!".to_string())
    } else {
        Ok(numerator / denominator)
    }
}

// if divide() returns an Err
// the ? operator will immediately return that Err
// thus preventing further execuution
fn calculate_average(a: f64, b: f64, c: f64) -> Result<f64, String> {
    let sum = a + b + c;
    let average = divide(sum, 3.0)?; // The '?' here propagates the error from 'divide'
    Ok(average)
}

fn main() {

    match calculate_average(10.0, 20.0, 0.0) {
        Ok(avg) => println!("Average: {}", avg),
        Err(e) => println!("Error: {}", e),
    }
}
```


## Turbofish operator `::<>`

Indicate the type information on the righthand side of the expression

```rust
let bytes = file.bytes().take(num_bytes).collect::<Result<Vec<_>, _>>();
```

Often it's just a matter of style.