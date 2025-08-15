---
id: Unbuffered channel
aliases:
  - Unbuffered channel
tags: []
---

# Unbuffered channel

## Blocking

A send operation on an unbuffered channel _blocks the sending goroutine until another goroutine executes a corresponding receive on the same channel_

Conversely, the receiving goroutine can be blocked if a receive operation happens first, and it will be unblocked when another goroutine performs a send operation.
