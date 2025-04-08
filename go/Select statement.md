---
id: Select statement
aliases: []
tags: []
---

Tags: #review #go #programming

Allow a goroutine to _wait on multiple communication operations_. This is useful when we have to _handle multiple concurrent operations_ and _prevent blocking behavior_

`select` blocks all cases until one of the cases can run, then it executes that case. If multiple cases are ready, it will choose a random case.

`select` is the reason why channels and goroutines are built into the language

```go
select {
    case msg1 := <-channel1:
    // Handle data from channel1
    case msg2 := <-channel2:
    // Handle data from channel2
    case channel3 <- data:
    // Send data to channel3
    default:
    // Optional: handle case when no channel is ready
}
```

Think of `select` like a `switch` with the catch as **each case is a communication**

## Timeout using select

`time.After` returns a channel that _blocks for the specified duration_. After that, the channel delivers ONCE
