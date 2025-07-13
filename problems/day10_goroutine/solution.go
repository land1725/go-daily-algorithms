package day10_goroutine

import (
	"fmt"
	"sync"
	"time"
)

func PrintOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	defer func() {
		fmt.Printf("PrintOdd elapsed: %v\n", time.Since(start))
	}()
	for i := 1; i < 10; i = i + 2 {
		fmt.Printf("%v ", i)
	}
}

func PrintEven(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	defer func() {
		fmt.Printf("PrintEven elapsed: %v\n", time.Since(start))
	}()
	for i := 0; i < 10; i = i + 2 {
		fmt.Printf("%v ", i)
	}
}
