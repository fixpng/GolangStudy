package main

import "fmt"

func print(n string) {
	fmt.Println(n)
}

func getData(s string) string {
	fmt.Println("getDate: ", s)
	return s
}

func Defer() {

	// 函数内异常捕获
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	n := 0
	fmt.Println(3 / n)
}

func main() {
	fmt.Println("1")

	defer print("2")          //defer 后面跟函数或方法
	defer print(getData("2")) //defer 后面跟函数或方法

	fmt.Println("3")

	Defer()
}
