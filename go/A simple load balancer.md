```go
// Simple load balancer

type Request struct {
	fn func() int // Operation to perform
	c  chan int   // Channel to return the result
}

type Worker struct {
	requests chan Request // Work to do (buffered channel)
	pending int
	index int
}

func requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Fake loading
		Sleep(rand.Int63n(nWorker * 2))
		work <- Request(workFn, c)
		result := <- c
		furtherProcess(result)
	}
}

// Send request to most lightly loaded worker
func (w *Worker) work(done chan *Worker) {
	// We could run the loop as a goroutine
	for {
		req := <-w.requests 	
		req.c <- req.fn()
		done  <- w
	}
}

type Pool []*Worker

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker
	w := heap.Pop(&b.Pool).(*Worker)
	// Send it the task
	w.requests <- req
	w.pending++
	// Put it on the heap
	heap.Push(&b.Pool, w)
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
			case req := <- work:
				b.dispatch(req) // Give the work to another worker
			case w := <- b.done:
				b.completed(w)		
		}
	}
}

// Job is completed, update heap
func (b *Balancer) completed(w *Worker) {
	w.pending--
	heap.Remove(&b.Pool, w.inde)
	heap.Push(&b.Pool, w) // Return it to its proper place	
}
 ```