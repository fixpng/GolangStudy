package main

import "fmt"

func main() {
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	// 自动推断长度
	var arr2 = [...]int{2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(arr3)

	arr = [...]int{9, 3, 2}
	//1.索引
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 二维数组
	var t = [3][4]int{
		{1, 2, 3, 4},
		{2, 4, 5, 6},
		{3, 4, 5, 6},
	}
	fmt.Println(t)

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			fmt.Printf("t[%v][%v]=%v\n", i, j, t[i][j])
		}
	}

	//三维数组
	var tt = [2][3][4]int{
		{
			{1, 2, 3, 4},
			{2, 4, 5, 6},
			{3, 4, 5, 6},
		}, {
			{1, 2, 3, 4},
			{2, 4, 5, 6},
			{3, 4, 5, 6}},
	}

	fmt.Println(tt)
}
