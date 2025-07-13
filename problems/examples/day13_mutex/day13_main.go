package main

import (
	"fmt"
	"github.com/land1725/go-daily-algorithms/problems/day13_mutex"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	safeCount := day13_mutex.SafeCounter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				safeCount.Inc()
			}
		}()
	}
	wg.Wait()
	safeCount.Value()

	automicCounter := day13_mutex.AtomicCounter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				automicCounter.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("safeCount: %d, automicCount: %d", safeCount.Value(), automicCounter.Value())

}
