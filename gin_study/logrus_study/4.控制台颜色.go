package main

import "fmt"

const (
	cBlack  = 0
	cRed    = 1
	cGreen  = 2
	cYellow = 3
	cBlue   = 4
	cPurple = 5
	cCyan   = 6
	cGray   = 7
)

func PrintColor(colorCode int, text string, isBackeGround bool) {
	if isBackeGround {
		fmt.Printf("\033[4%dm%s\033[0m\n", colorCode, text)
	} else {
		fmt.Printf("\033[3%dm%s\033[0m\n", colorCode, text)
	}
}

func PrintColor0() {
	// 前景色
	fmt.Println("\033[30m 黑色 \033[0m")
	fmt.Println("\033[31m 红色 \033[0m")
	fmt.Println("\033[32m 绿色 \033[0m")
	fmt.Println("\033[33m 黄色 \033[0m")
	fmt.Println("\033[34m 蓝色 \033[0m")
	fmt.Println("\033[35m 紫色 \033[0m")
	fmt.Println("\033[36m 青色 \033[0m")
	fmt.Println("\033[37m 灰色 \033[0m")
	// 背景色
	fmt.Println("\033[40m 黑色 \033[0m")
	fmt.Println("\033[41m 红色 \033[0m")
	fmt.Println("\033[42m 绿色 \033[0m")
	fmt.Println("\033[43m 黄色 \033[0m")
	fmt.Println("\033[44m 蓝色 \033[0m")
	fmt.Println("\033[45m 紫色 \033[0m")
	fmt.Println("\033[46m 青色 \033[0m")
	fmt.Println("\033[47m 灰色 \033[0m")
}

func main() {
	PrintColor(cBlue, "蓝色", false)
	PrintColor(cBlue, "蓝色", true)

	fmt.Println("123456789啦啦啦啦啦")
	fmt.Println("\033[31m 你好，这里的颜色是红色 \033[0m")
	PrintColor0()
}
