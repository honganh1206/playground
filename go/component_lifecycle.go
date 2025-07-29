package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jacoelho/component"
)

type Lifecycle interface {
	Start(context.Context) error
	Stop(context.Context) error
}

type Logger struct{}

func (l *Logger) Log(message string)              { fmt.Println(message) }
func (l *Logger) Start(ctx context.Context) error { l.Log("Logger started"); return nil }
func (l *Logger) Stop(ctx context.Context) error  { l.Log("Logger stopped"); return nil }

type MainService struct {
	logger *Logger
}

func (s *MainService) Start(ctx context.Context) error {
	s.logger.Log("MainService started")
	return nil
}

func (s *MainService) Stop(ctx context.Context) error {
	s.logger.Log("MainService stopped")
	return nil
}

// Problem: Lifecycle logic is scattered here
// since main goroutine will finish before workers have time to clean up
func Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	httpCtx, httpCancel := context.WithCancel(context.Background())
	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
		BaseContext: func(l net.Listener) context.Context {
			return httpCtx
		},
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// If we run the server on the main goroutine,
		// the server would block the main goroutine
		// and we cannot run other tasks like worker goroutine
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server failed: %v", err)
		}
	}()

	workerCtx, workerCancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Listen indefinitely until there is a element in Done channel
		// Think of it like a stop signal
		for {
			select {
			case <-workerCtx.Done():
				log.Println("worker stopped")
				return
			default:
				// Simulate queue processing
				log.Println("processing queue item")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	log.Println("server and worker started")
	<-ctx.Done()
	log.Println("shutdown signal received")

	// Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("initiating graceful shutdown")
	// Shutdown waits for in-flight HTTP requests to finish
	// But will not wait for background goroutines to finish (workers)
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}

	// Use these immediately instead of deferring them
	// i.e., waiting for main() to return
	// to signal goroutines using these contexts to stop ASAP
	// and thus more predictable shutdown process
	httpCancel()
	workerCancel()
	wg.Wait()
	log.Println("shutdown complete")

}

// What is better: Reduce manual writings of start/stop logic
// and enfore correct ordering via explicit dependency declarations
func BetterRun() {
	// Central registry tracking components and dependencies
	// Orchestrating start and stop of components
	sys := new(component.System)
	ctx := context.Background()

	// Identifiers for registered components
	var (
		loggerKey  component.Key[*Logger]
		serviceKey component.Key[*MainService]
	)

	// Logger component will start first
	if err := component.Provide(sys, loggerKey, func(_ *component.System) (*Logger, error) {
		return &Logger{}, nil
	}); err != nil {
		log.Fatalf("failed to provide logger: %v", err)
	}

	// MainService depends on Logger
	// and will start second
	if err := component.Provide(sys, serviceKey, func(s *component.System) (*MainService, error) {
		logger, err := component.Get(s, loggerKey) // Get dependency
		if err != nil {
			return nil, err
		}
		return &MainService{logger: logger}, nil
	}, loggerKey); err != nil { // Declare dependency on loggerKey
		log.Fatalf("Failed to provide main service: %v", err)
	}

	fmt.Println("Starting system...")
	startCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	// Start components in topological order
	if err := sys.Start(startCtx); err != nil {
		log.Fatalf("System start failed: %v", err)
	}

	fmt.Println("System is UP.")

	fmt.Println("Stopping system...")
	stopCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := sys.Stop(stopCtx); err != nil {
		log.Printf("System stop encountered errors: %v", err)
	}
	fmt.Println("System shut down.")
}

func main() {
	// Run()
	BetterRun()
}
