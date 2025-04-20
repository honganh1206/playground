---
id: Concurrency is NOT Parallelism
aliases:
  - Concurrency is NOT Parallelism
tags: []
---

# Concurrency is NOT Parallelism

[Ref](https://go.dev/s/concurrency-is-not-parallelism)

They are not the same thing

Concurrency - Composition of independently executed processes (Dealing with a lot of things at once) # Parallelism - Simultaneous execution of multiple things (Doing a lot of things at once)

Goroutines are cheaper than threads resource-wise


[[A simple load balancer]]

[[Replicated database  and minimize latency]]

Concurrency enables parallelism - The tools of concurrency makes it trivial to build safe and scalable parallel design

Concurrency simplifies synchronization - No explicit synchronization needed like Java with `Thread` and `volatile` 