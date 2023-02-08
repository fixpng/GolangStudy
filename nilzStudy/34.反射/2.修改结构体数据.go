package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
	Name string `json:"name" feng:"name_xxx"`
	Age  int    `json:"age" feng:"age_xxx"`
}

func FPrint1(inter interface{}) {
	v := reflect.ValueOf(inter)
	e := v.Elem() // 必须用这个
	e.FieldByName("Name").SetString("茜茜知道")
}

func main() {
	user := User1{"茜茜", 2}
	fmt.Println(user)
	FPrint1(&user) // 必须传指针
	fmt.Println(user)
}
