package main

import "log"

// This file does not actually run any code

// Embarrassingly parallel: Subproblems independent of each other.
// These are the easiest kind to implement, and they scale linearly with the amount of parallelism.
func makeThumbnails1(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// First attempt to execute all operations in parallel
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		// Start all goroutines, one per filename
		// but does not wait for them to finish
		go ImageFile(f) // Ignoring errors
	}
}
