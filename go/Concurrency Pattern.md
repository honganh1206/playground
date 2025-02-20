---
id: Concurrency Pattern
aliases: []
tags: []
---

# Concurrency Pattern in Go

[Ref](https://go.dev/blog/pipelines)

A way to construct streaming data pipelines which utilize I/O and multiple CPUs

## Pipelines in Go

A series of stages connected by [channels](./Channels.md). Each stage is a
**group of goroutines** running the same function

Things happen during each stage:

1. Receive values from upstream via inbound channels
2. Process the values
3. Send values downstream via outbound channels

The 1st stage only has inbound channels (source/producer), while the 2nd stage
has only outbound channels (sink/consumer)

[[Fan in and fan out]]

## Stopping short
