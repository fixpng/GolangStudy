package main

import "fmt"

func main11() {
	// 与 或 非
	fmt.Println(2 > 1 || 1 > 2)
	fmt.Println(2 > 1 && 1 > 2)
	fmt.Println(!(2 > 1))
	fmt.Println("A" > "a") //字符串就是比ascii值

	// 逻辑短路
	// && 前面如果是false，后面就不判断了，直接为false
	a, b := 10, 0
	if a < 0 && a/b > 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("xx")
	}

	// || 前面如果是true，后面就不判断了，直接为true
	if a > 0 || a/b > 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("xx")
	}

}
