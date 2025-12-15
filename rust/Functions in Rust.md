## Closure

```rust
// || instead of () for input variables
// {} for single-line expression (optional)
let closure_annotated = |i: i32| -> i32 { i + outer_var };
let closure_inferred  = |i     |          i + outer_var  ;
```