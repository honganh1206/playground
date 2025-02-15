Tags: #review #programming #golang

## What is `sync.Mutex`

- Short for _mutual exclusion_
- `sync.Mutex` set a lock around particular lines of code
- While one Goroutine holds the lock, all other Goroutines are prevented from executing any lines of code protected by the same mutex => Goroutines are forced to wait until the lock is yielded before they can proceed.

```go
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock() // Ensure the mutex will be unlocked
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey") // During the lock only this goroutine is run
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey")) // THe lock is now given to the main goroutine
}

```

> [!IMPORTANT]
> We hold locks for the **shortest time possible** to maintain good performance while ensuring thread safety

> [!note]
>
> The same `mutex` global variable can be used in **multiple places** throughout your code so long as it is the same mutex

Exercise: Web Crawler

```go
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Track visited URLs
type SafeVisited struct {
	mu sync.Mutex
	visited map[string]bool
}

func (v *SafeVisited) IsVisited(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.visited[url] {
		return false
	}
	v.visited[url] = true
	return true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, visited *SafeVisited, wg *sync.WaitGroup) {
	defer wg.Done() // Executed when we are done with EACH goroutine adding an URL to the map

	if depth <= 0 {
		return
	}

	if !visited.IsVisited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// Crawl URLs in parallel
	for _, u := range urls {
		wg.Add(1) // Incrementing the counter for each newly spawned gorountine
		go Crawl(u, depth-1, fetcher, visited, wg)
	}
	return
}

func main() {
	visited := &SafeVisited{visited: make(map[string]bool)}

	var wg sync.WaitGroup

	wg.Add(1) // Notify that we now track the 1st goroutine

	go Crawl("https://golang.org/", 4, fetcher, visited, &wg) // Ensure the same wg object is passed in

	wg.Wait() // wait for all goroutines to complete
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}


```
