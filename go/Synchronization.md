---
id: Synchronization
aliases:
  - Synchronization
tags: []
---

# Synchronization

It is like how [[Channels]] works: When the main function executes `<-c`, it will wait for a value
to be sent

When the invoked function executes `c <- value`, it will wait for a receiver to be ready

A sender and receiver must **both be ready** to play their part. Otherwise we wait until they are
