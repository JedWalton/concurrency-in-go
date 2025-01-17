package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Number of workers
//const workerCount = 10

func main() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Ideally allocate the available number of CPUs for the worker count
	workerCount := runtime.NumCPU()

	dataCh := make(chan int)
	var wg sync.WaitGroup
	// Launch a fixed number of workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for val := range dataCh {
				// Simulate some trivial processing
				_ = val * val
			}
		}(i)
	}
	// Send some data to the channel
	for i := 0; i < 1000; i++ {
		dataCh <- i
	}

	fmt.Printf("Active Goroutines: %d\n", runtime.NumGoroutine())

	close(dataCh)
	// Wait for workers to complete
	wg.Wait()
}
