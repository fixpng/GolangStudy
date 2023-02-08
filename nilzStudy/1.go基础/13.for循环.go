package __go基础

import "fmt"

func main13() {

	for i := 0; i < 10; i++ {
		fmt.Printf("打印, %d\n", i)
	}

	i := 1
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}
}
