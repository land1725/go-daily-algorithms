package main

import (
	"fmt"
	"github.com/land1725/go-daily-algorithms/problems/day10_goroutine"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go day10_goroutine.PrintOdd(&wg)
	go day10_goroutine.PrintEven(&wg)
	fmt.Printf("exit")
	wg.Wait()
}
