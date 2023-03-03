package main

import "fmt"

type msgType uint16
type mt = uint16

func main() {
	//自定义类型
	var successMsg msgType = 1000
	var errorMsg msgType = 1210
	var num uint16 = 23

	fmt.Printf("successMsg:%v,%T\n", successMsg, successMsg)
	fmt.Printf("errorMsg:%v,%T\n", errorMsg, errorMsg)
	fmt.Printf("num:%v,%T\n", num, num)

	successMsg = msgType(num)
	fmt.Printf("successMsg:%v,%T\n", successMsg, successMsg)

	//类型别名
	var u uint8 = 97
	var b byte = 97
	fmt.Printf("successMsg:%v,%T\n", u, u)
	fmt.Printf("successMsg:%v,%T\n", b, b)

	var r rune = '中'
	var ui int32 = '中'
	fmt.Printf("successMsg:%v,%T\n", r, r)
	fmt.Printf("successMsg:%v,%T\n", ui, ui)

	var successMsg2 mt = 1005
	var uin uint16 = 1006
	fmt.Printf("successMsg:%v,%T\n", successMsg2, successMsg2)
	fmt.Printf("successMsg:%v,%T\n", uin, uin)

}
