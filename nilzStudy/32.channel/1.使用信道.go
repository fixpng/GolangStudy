package main

import "fmt"

func main() {
	//var ch chan string = make(chan string, 2)
	var ch chan interface{} = make(chan interface{}, 2)

	ch <- "茜茜" // 写入数据到信道
	ch <- 3

	s := <-ch
	fmt.Println(s)
	fmt.Println(<-ch)

	ch <- "茜茜" // 写入数据到信道
	ch <- "知道"

	ss, ok := <-ch
	fmt.Println(ss, ok)
	fmt.Println(<-ch)

	close(ch)
}
