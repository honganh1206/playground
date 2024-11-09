
## What is `sync.Mutex`

- `sync.Mutex` set a lock around particular lines of code
- While one goroutine holds the lock, all other goroutines are prevented from executing any lines of code protected by the same mutex => Goroutines are forced to wait until the lock is yielded before they can proceed.

> [!note]
> 
> The same `mutex` global variable can be used in **multiple places** throughout your code so long as it is the same mutex


### Read Write Mutexes

- `sync.RWMutex` allows **any** number of readers to hold the lock or **one** writer => More efficient than using a full mutex