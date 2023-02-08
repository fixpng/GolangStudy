package __go基础

import (
	"fmt"
)

// 批量声明
var (
	m78 string
	m87 int = 10
	m90 bool
)

func main4() {

	var name string
	var age int
	var score float32
	var isSuper bool
	// 针对基础类型 如果声明之后不赋值，会有默认值
	fmt.Println(name, age, score, isSuper)

	name = "Nilz"
	fmt.Println(name, age, score, isSuper)

	// 自动推断类型
	var userName = 244
	fmt.Println(userName)

	// 声明赋值
	userAddr := "广州"
	fmt.Println(userAddr)

	fmt.Println(m78, m87, m90)

}
