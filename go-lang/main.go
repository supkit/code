package main

import (
	"fmt"
	"sync"
)

// https://www.mogublog.net/post/2120.html
var wg sync.WaitGroup

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("write data: ", i)
	}
}

func read(ch chan int) {
	for {
		fmt.Println("read data: ", <-ch)
		wg.Done()
	}
}

func main() {
	ch := make(chan int, 5)
	wg.Add(10)

	go write(ch)
	go read(ch)

	wg.Wait()
}
