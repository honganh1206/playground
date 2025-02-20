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
	mu      sync.Mutex
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
