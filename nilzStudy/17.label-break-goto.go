package main

import "fmt"

func main17() {
	//start:
	//	for i := 1; i <= 5; i++ {
	//		for j := 1; j <= 6; j++ {
	//			fmt.Printf("*")
	//			if i == 3 && j == 2 {
	//				break start
	//			}
	//		}
	//		fmt.Println()
	//
	//	}

	//	for i := 1; i <= 5; i++ {
	//		for j := 1; j <= 6; j++ {
	//			fmt.Printf("*")
	//			if i == 3 && j == 2 {
	//				goto start
	//			}
	//		}
	//		fmt.Println()
	//
	//	}
	//start:

	var (
		username string
	)
start:
	fmt.Println("请输入用户名")
	fmt.Scanln(&username)
	if username != "nlz" {
		goto start
	}

}
