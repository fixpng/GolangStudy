package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))

	sum := add(2, 4)
	fmt.Println(sum)
}
