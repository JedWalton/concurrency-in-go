package main

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runOptimized()
	}
}

func runOptimized() {
	dataCh := make(chan int)
	var wg sync.WaitGroup

	// Allocate the available number of CPUs
	workerCount := runtime.NumCPU()

	// Create a fixed number of workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range dataCh {
				_ = val * val
			}
		}()
	}
	// Send some data to the channel
	go func() {
		for i := 0; i < 1000; i++ {
			dataCh <- i
		}
		close(dataCh)
	}()
	// Wait for all workers to finish
	wg.Wait()
}
