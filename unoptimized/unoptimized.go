package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// We'll create a channel that we'll push a bunch of integers onto.
	dataCh := make(chan int)
	// Create a WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup
	// Spin up a large number of goroutines (naive approach).
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range dataCh {
				// Simulate some trivial processing work
				_ = val * val
			}
		}()
	}
	// Send some data to the channel
	for i := 0; i < 1000; i++ {
		dataCh <- i
	}

	fmt.Printf("Active Goroutines: %d\n", runtime.NumGoroutine())

	// Close the channel to signal goroutines to stop
	close(dataCh)
	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Done with unoptimized version")
}
