---
id: Fan in and fan out
aliases: []
tags: []
---

Multiple functions can read _from the same channel_ until that channel is closed
(Fan out). This is to _distribute work_ amongst a group of workers & parallelize
CPU use and I/O

A function processes multiple data streams. Once all are read through, their
channels are combined instantly onto a single channel that closes once all
streams have closed.

[Fan in code example](./fan_in.go)
