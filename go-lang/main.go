package main

import "fmt"

// 管道学习文章 http://c.biancheng.net/view/97.html

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println(1212)
		ch <- 0
	}()
	go func() {
		fmt.Println(13123)
		ch <- 0
	}()
	<-ch
	fmt.Println(1)
}
