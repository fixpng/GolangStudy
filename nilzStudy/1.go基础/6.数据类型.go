package __go基础

import (
	"fmt"
	"math"
)

func main6() {
	var i8 int8
	i8 = 127

	fmt.Println(i8)
	fmt.Println(math.Pow(2, 3))

	var b uint8 = 0b001000000
	fmt.Printf("0b%b,%d\n", b, b)

	var c uint8 = 0xa1 //10 * 16 + 1 * 1
	fmt.Printf("0b%x,%d\n", c, c)

	// 浮点数
	var f float32 = 3.1415926
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)
}
