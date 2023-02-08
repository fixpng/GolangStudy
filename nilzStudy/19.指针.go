package main

import "fmt"

func setName(name string) {
	fmt.Printf("函数内部：%p\n", &name)
}

func setName2(name *string) {
	// 字符串变量name映射的地址
	*name = "我是"
	fmt.Printf("函数内部：%p\n", name)
}

func main() {

	name := "茜茜西"
	fmt.Printf("函数外部：%p\n", &name)
	setName(name)
	fmt.Println(name)

	setName2(&name)

}
