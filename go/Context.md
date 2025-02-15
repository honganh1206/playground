# Context library

A library to manage **cancellation, deadlines and request-scoped values**

```go
import "context"

// Basic context types
context.Context        // Interface
context.Background()   // Root context
context.TODO()        // Placeholder context
```

## Context creation

```go
// Creating contexts
ctx := context.Background()
ctx := context.TODO()

// With cancellation
ctx, cancel := context.WithCancel(parentCtx)
defer cancel() // Always remember to cancel

// With timeout
ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
defer cancel()

// With deadline
ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(5*time.Second))
defer cancel()

// With value
ctx = context.WithValue(parentCtx, key, value)
```

## Common patterns

A. HTTP Server Example:

```go
func HandleRequest(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    // Use context for the request lifetime
    select {
    case <-ctx.Done(): // Check long-running operation
        return
    case result := <-doWork(ctx):
        fmt.Fprintf(w, result)
    }
}
```

B. Database Query Example:

```go
func QueryWithTimeout(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()

    return db.QueryRowContext(ctx, "SELECT ...").Scan(&result)
}
```

## Best practices

• Context Propagation:

- Pass context as first parameter
- Don't store contexts in structs
- Always pass contexts explicitly

• Cancellation:

- Always call cancel() functions
- Use defer for cleanup
- Check `ctx.Done()` in long-running operations

• Value Management:

- Use context values sparingly
- Keep values request-scoped
- Use strong types for context keys

## Common use cases

A. Request Tracing:

```go
func AddTraceID(ctx context.Context, traceID string) context.Context {
    return context.WithValue(ctx, traceIDKey, traceID)
}
```

B. Timeout Management:

```go
func TimeoutOperation(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
    defer cancel()

    select {
    case <-ctx.Done(): // In case of a long-running operation
        return ctx.Err()
    case result := <-operation():
        return result
    }
}
```

## Important considerations

• Thread Safety:

- Context is safe for concurrent use
- Values are immutable once stored
- New contexts create new value trees

• Error Handling:

```go
if err := ctx.Err(); err != nil {
    switch err {
    case context.Canceled:
        // Handle cancellation
    case context.DeadlineExceeded:
        // Handle timeout
    }
}
```

## Anti-patterns to avoid

• Storing contexts in structs
• Using context for optional parameters
• Passing nil contexts
• Failing to cancel contexts
• Using context.Background() in request handlers
