package main

import "fmt"

// Sum 函数接收多个参数
func Sum(nums ...int) (sum int) {
	for i, v := range nums {
		fmt.Println(i)
		sum += v
	}
	return
}

func main() {
	var s []int = []int{1, 2, 3, 4}
	fmt.Println(s, len(s), cap(s))
	//切片是怎样扩容的
	s = append(s, 3, 4)
	fmt.Println(s, len(s), cap(s))

	//复制
	var s1 = make([]int, 8, 8)
	copy(s1, s)
	fmt.Println(s1)

	// string 与 []byte
	str := "hello 世界"
	fmt.Printf("[]byte(str)=%s,%v\n", []byte(str), []byte(str))
	fmt.Println(len(str))

	for i, v := range str {
		fmt.Printf("str[%d]=%c\n", i, v)
	}

	fmt.Println(Sum(2, 4, 6, 87, 10))
	fmt.Println(Sum(87, 10))
}
