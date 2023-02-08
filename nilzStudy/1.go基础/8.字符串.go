package __go基础

import "fmt"

func main8() {
	var c0 uint8 = 65
	var c1 byte = 65
	var c2 = 'a'
	var c3 rune = '中'

	fmt.Printf("c0 码值：%v, 字符为：%c，类型是%T。\n", c0, c0, c0)
	fmt.Println(c0 == c1)
	fmt.Printf("c1 码值：%v, 字符为：%c，类型是%T。\n", c1, c1, c1)
	fmt.Printf("c2 码值：%v, 字符为：%c，类型是%T。\n", c2, c2, c2)
	fmt.Printf("c3 码值：%v, 字符为：%c，类型是%T。\n", c3, c3, c3)

	var s1 string = `倪半仙\n知道`
	var s2 string = "倪半仙\n知道"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s1))

}
