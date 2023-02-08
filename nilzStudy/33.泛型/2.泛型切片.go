package main

import "fmt"

type mySlice[T int | string] []T

// 泛型函数
func PrintIntTypeSlice2[T int | int64 | string](slice []T) {
	fmt.Printf("%T\n", slice)
	for _, v := range slice {
		fmt.Printf("%T  %v\n", v, v)
	}
}
func main() {
	v1 := mySlice[int]{1, 2, 3, 4, 5}
	PrintIntTypeSlice2(v1)
	v2 := mySlice[string]{"1, 2, 3, 4, 5", "hello"}
	PrintIntTypeSlice2(v2)
}
