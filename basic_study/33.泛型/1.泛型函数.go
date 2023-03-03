package main

import "fmt"

// 普通区
func PrintIntSlice(slice []int) {
	for _, i2 := range slice {
		fmt.Printf("%T  %v\n", i2, i2)
	}
}
func PrintInt64Slice(slice []int64) {
	for _, i2 := range slice {
		fmt.Printf("%T  %v\n", i2, i2)
	}
}

// 泛型函数
func PrintIntTypeSlice[T int | int64 | string](slice []T) {
	fmt.Printf("%T\n", slice)
	for _, v := range slice {
		fmt.Printf("%T  %v\n", v, v)
	}
}

func main() {
	// 普通区
	/*
		PrintIntSlice([]int{1, 2, 3, 4, 5})
		var int64Slice []int64 = []int64{4, 5, 7}

		PrintInt64Slice(int64Slice)

		var intSlice []int
		for _, v := range int64Slice {
			intSlice = append(intSlice, int(v))
		}
		PrintIntSlice(intSlice)
	*/

	PrintIntTypeSlice([]int{1, 2, 3, 4, 5})
	PrintIntTypeSlice([]int64{1, 2, 3, 4, 5})
	PrintIntTypeSlice([]string{"1, 2, 3, 4, 5", "hello"})
}
