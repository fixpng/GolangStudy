package main

import "fmt"

type data interface {
}

type Dog struct {
	Name string
}

func Print(d data) {
	fmt.Println(d)
}

func PrintNum(num interface{}) {
	if value, ok := num.(int); ok {
		fmt.Println("int", value)
	} else if value, ok := num.(string); ok {
		fmt.Println("string", value)
	} else {
		fmt.Println("不明类型")
	}

	switch num.(type) {
	case int:
		fmt.Println("int", num)
	case string:
		fmt.Println("string", num)
	}

}

func main() {
	d := Dog{"小黑"}
	Print(d)
	Print(12)
	Print("124")
	Print(true)
	Print([]int{1, 2, 3})
	PrintNum(34238888888888888.45)
	PrintNum(342)

}
