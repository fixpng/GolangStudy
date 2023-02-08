package main

import (
	"GO20230123/nilzStudy/version"
	"fmt"
)

// 批量声明
// iota 一个括号里的层数
const (
	STATE int = 1
	a     int = iota
	b     int = 1
	c     int = iota
)

func main5() {
	fmt.Println(STATE)
	fmt.Println(a, b, c)

	fmt.Println(version.Version)
	//fmt.Println(version.STATE)
}
