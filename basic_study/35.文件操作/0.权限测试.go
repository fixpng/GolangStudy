package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.FileMode(777), 777)
	fmt.Println(os.FileMode(0777), 0777)

	um, _ := strconv.ParseInt(strconv.Itoa(777), 8, 0)
	fmt.Println(os.FileMode(777), 777)
	fmt.Println(os.FileMode(0777), 0777)
	fmt.Println(os.FileMode(um), um)
}
