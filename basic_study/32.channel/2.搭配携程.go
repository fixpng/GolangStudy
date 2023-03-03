package main

import "fmt"

var ch chan int = make(chan int)

func insertNum() {

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	go insertNum()

	for {
		i, ok := <-ch

		fmt.Println(i, ok)
	}
}
