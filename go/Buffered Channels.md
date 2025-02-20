---
id: Buffered Channels
aliases: []
tags: []
---

Tags: #review #golang #programming

Go can have buffered channels (channels holding a certain number of values).

```go
func main() {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 2 // Adding this leads to a deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

```

|                    | Buffered                               | Un-buffered                                  |
| ------------------ | -------------------------------------- | -------------------------------------------- |
| Sender -> Receiver | Sender blocks if the buffer is full    | Sender blocks until receiver ready to read   |
| Receiver -> Sender | Receiver blocks if the buffer is empty | Receiver blocks until there is value to read |

> [!IMPORTANT] Buffering removes [synchronization](./Synchronization.md). Buffering makes channels
> more like Erlang's mailboxes?
