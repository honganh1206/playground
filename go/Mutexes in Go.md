# Mutexes

- Short for _mutual exclusion_
- `sync.Mutex` set a lock around particular lines of code
- While one Goroutine holds the lock, all other Goroutines are prevented from executing any lines of code protected by the same mutex => Goroutines are forced to wait until the lock is yielded before they can proceed.

```go
// SafeCounter is safe to use concurrently.
```

> [!IMPORTANT]
> We hold locks for the **shortest time possible** to maintain good performance while ensuring thread safety

> [!note]
>
> The same `mutex` global variable can be used in **multiple places** throughout your code so long as it is the same mutex
