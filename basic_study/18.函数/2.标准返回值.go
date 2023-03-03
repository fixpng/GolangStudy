package main

import "fmt"

func getAddDiff(a int, b int) (int, int) {
	return a + b, a - b
}
func main() {
	fmt.Println(getAddDiff(1, 2))

	x, y := getAddDiff(2, 4)
	fmt.Println(x, y)
}
