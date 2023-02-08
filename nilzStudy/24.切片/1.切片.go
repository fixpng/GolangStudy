package main

import "fmt"

func main() {
	var arr = [5]int{3, 4, 5, 6, 7}
	var O = arr[0]
	fmt.Println(O)

	fmt.Println(arr, len(arr))

	fmt.Println(arr[0:2], arr[0:], arr[0:len(arr)])
}
