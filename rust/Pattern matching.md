Pattern matching means:
- **Check the label**
- **If it matches**, pull out the contents

Imagine you have **mystery boxes**.  
Each box has a **label** and maybe **something inside**.

Examples:

- A box labeled **“OK”** with a number inside
    
- A box labeled **“ERR”** with a message inside
    

You don’t open boxes randomly.  
You say:

> “If the box says **OK**, open it and give me what’s inside.”

That rule is **pattern matching**.

```rust
if let Ok(val) = num.parse() {
    // use val
}
// “Try parsing.  
// If the result is an **OK box**, take what’s inside and call it `val`.  
// Otherwise, do nothing.”
```