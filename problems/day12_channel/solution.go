package day12_channel

import "fmt"

func ReadChan(ch <-chan int) {
	for v := range ch {
		fmt.Printf("读取:%d\n", v)
	}
}
func WriteChan(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 100; i++ {
		fmt.Printf("写入:%d\n", i)
		ch <- i
	}
}
