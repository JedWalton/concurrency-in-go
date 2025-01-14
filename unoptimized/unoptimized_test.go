package main

import (
	"testing"
)

func BenchmarkUnoptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runUnoptimized()
	}
}

func runUnoptimized() {
	dataCh := make(chan int)
	doneCh := make(chan struct{})
	go func() {
		for i := 0; i < 1000; i++ {
			dataCh <- i
		}
		close(dataCh)
	}()
	// Spin up a large number of goroutines
	for i := 0; i < 100000; i++ {
		go func() {
			for val := range dataCh {
				_ = val * val
			}
		}()
	}
	// No WaitGroup here for simplicity in the benchmark,
	// but we'll simulate waiting a bit.
	go func() {
		<-dataCh
		close(doneCh)
	}()
	<-doneCh
}
