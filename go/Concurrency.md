---
id: Concurrency in Go
aliases: []
tags: []
---

## Refs

[[Goroutines]]

[[Channels]]

[[Mutexes]]

[[Concurrency Pattern]]

[Slide by Rob Pike](https://go.dev/talks/2012/concurrency.slide)

## What is concurrency?

Concurrency is the **composition** of _independently executing computations_

Concurrency mirrors a complex world of interacting, independently behaving pieces

[[Concurrency is NOT Parallelism]], but it **enables** parallelism. If we have only one processor, we can enable concurrency but not parallelism

A well-written concurrent program can run _efficiently_ in parallel on a multiprocessor

Some languages sharing the concurrency feature are Occam, Erlang, Newsqueak, etc.,

## Discussions on concurrency

When we say _x happens before y_, we don't mean merely that _x occurs earlier than y in time_, but we mean that x is **guaranteed** to do so only (since multicore CPU can have its instructions reordered by compilers). In this sense, we can safely read results from x in without worrying about inconsistent data.

When x neither happens before nor after y, we say **x is concurrent with y**. In this sense, we cannot assume anything about their ordering.

[[Concurrency in Go]]

[[Communication]]

[[Synchronization]]

[[Multiplexing]]

[[Writing system software with Go]]
