---
id: Channels
aliases: []
tags: []
---

Tags: #review #programming #golang

# Channels in Go

A **typed conduit** through which we _send and receive values_. It is a way for goroutines to [communicate](./Communication.md)

```go
ch := make(chan int) // Create a channel
ch <- v // Send v to channel ch
v := <-ch // Receive value from ch and assign value to x
```

> [!important] Important Sends and receives are blocked until the other side is ready.
> This allows goroutines to sync without explicit locks or condition variables.

As with maps, a channel is a _reference_ to the data structure created by `make`. This means when we copy a channel or pass one as an argument, we are copying a _reference_ of it and the caller + callee refer to the same data structure.

Two channels can be compared using `==`, and the result is `true` if both refer to the same channel data structure

A channel has _two principal operations: send and receive_. A send statement transmits a value from one goroutine through a channel to another goroutine which execute a corresponding receive expression

```go
ch <- x // Send statement
x = <- ch // Receive expression
<- ch // Result of receive expression is discarded
```

We can have _3 types of channels_: Bidirectional, receive-only and send-only

```go
chan time.Time // Bidirectional (can both send and receive)
<-chan time.Time // Receive-only
chan<- time.Time // Send-only
```

A channel created with `make(chan int)` only is an [[Unbuffered channel]] . If we input an optional second argument, we have [Buffered Channels](./Buffered%20Channels.md)

## Range and Close

A sender can use the keyword `close` to indicate that no more values will be sent, and receivers can test if the channel has been closed by using a second boolean parameter (See `fib.go`)

After a channel has been closed, any further send operations on it will panic. After the closed channel is drained i.e., last sent element has been received, all subsequent receive operations will proceed but will yield a zero value.

[Select statement](./Select%20statement.md)

A channel both [communicates](./Communication.md) and [synchronizes](./Synchronization.md)

## Messages in channels

Messages sent over channels have two important aspects:

1. The value of the message
2. The moment at which the message occurs
