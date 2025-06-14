package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type CustomHandler struct {
	jobqueue chan string
}

func NewCustomHandler(jobqueue chan string) *CustomHandler {
	// You can check for wg == nil if feeling paranoid
	return &CustomHandler{jobqueue: jobqueue}
}

// Simulate 3 slow jobs e.g,, writing to s3, kafka, etc.
func slowJob1(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("starting job 1 for %s\n", name)
	time.Sleep(5 * time.Second)
	fmt.Printf("finished job 1 for %s\n", name)
}

func slowJob2(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("starting job 2 for %s\n", name)
	time.Sleep(4 * time.Second)
	fmt.Printf("finished job 2 for %s\n", name)
}

func slowJob3(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("starting job 3 for %s\n", name)
	time.Sleep(3 * time.Second)
	fmt.Printf("finished job 3 for %s\n", name)
}

// Implementation of http.Handler
func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobName := vars["jobName"]
	h.jobqueue <- jobName

	fmt.Fprintf(w, "job %s started", jobName)
}

// Stop writing new jobs and new connections
func consumer(ctx context.Context, jobqueue chan string, done chan any) {
	wg := &sync.WaitGroup{}

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("writing to done channel")
			done <- struct{}{}
			log.Println("Done, shutting down the consumer")
			return
		case job := <-jobqueue:
			wg.Add(3)
			go slowJob1(job, wg)
			go slowJob2(job, wg)
			go slowJob3(job, wg)

		}
	}
}

func main() {
	jobQueue := make(chan string, 10) // Buffered channel for better performance
	customHandler := NewCustomHandler(jobQueue)

	router := mux.NewRouter()
	router.Handle("/{jobName}", customHandler)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	// Separate contexts for consumer and server shutdown
	consumerCtx, cancelConsumer := context.WithCancel(context.Background())
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelShutdown()

	doneChan := make(chan any)

	// Start the consumer first
	go consumer(consumerCtx, jobQueue, doneChan)

	// Start the HTTP server
	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v\n", err)
		}
	}()

	// Wait for SIGTERM to be captured
	go func() {
		<-termChan
		log.Println("SIGTERM received. Shutdown process initiated")

		// Cancel the consumer context to stop processing new jobs
		cancelConsumer()

		// Shutdown the HTTP server
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server Shutdown Failed: %+v", err)
		}
		log.Println("HTTP server shut down")
	}()

	// Wait for the consumer's jobs to finish
	log.Println("Server running. Waiting for shutdown signal...")
	<-doneChan
	log.Println("All jobs completed. Shutting down.")
}
