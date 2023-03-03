package main

import "fmt"

type User1 struct {
	Name     string
	Age      int
	password string
}

func main() {
	var u1 User1 = User1{"绽放三", 21, "7asd5ff3"}
	fmt.Printf("u1:%#v,%T\n", u1, u1)
	fmt.Printf("u1:%v,%T\n", u1, u1)

	//第三种
	var u3 User1
	u3.Name = "zhangsan"
	u3.Age = 22
	u3.password = "#$%DGfr4g4"
	fmt.Printf("u3:%#v,%T\n", u3, u3)

	u4 := User1{"AA", 34, "GF345G4"}
	fmt.Printf("u4:%#v,%T\n", u4, u4)

	fmt.Println(u4.Name)
}
