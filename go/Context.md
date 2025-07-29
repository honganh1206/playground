---
id: Context
aliases: []
tags: []
---

## Definition

A signal/controller to manage *how long some work should continue*. A remote control to stop a goroutine

In Go, goroutines run independently, so we use `context` when we want to stop them early (cancel)/set a time limit (timeout)/pass request-scoped value

```go
// Example
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-time.After(3 * time.Second):
    fmt.Println("Finished work!")
case <-ctx.Done():
    fmt.Println("Too slow, timeout!")  // This prints after 2 seconds
}
```

![warning]
> Context cancellation only signals goroutines to stop. It does NOT wait for them to stop.

## Real-world analogy

Imagine you send someone on a task and give them a walkie-talkie (context). You can:

- Say “Stop!” anytime (cancel())
- Or let their timer beep after 2 minutes (WithTimeout)
- They keep working until they hear a stop signal `<-ctx.Done()`

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
