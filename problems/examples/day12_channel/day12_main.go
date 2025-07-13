package main

import (
	"github.com/land1725/go-daily-algorithms/problems/day12_channel"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		day12_channel.ReadChan(ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		day12_channel.WriteChan(ch)
	}()
	wg.Wait()

}
