package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" feng:"name_xxx"`
	Age  int    `json:"age" feng:"age_xxx"`
}

func Print(inter interface{}) {
	switch x := inter.(type) {
	case User:
		fmt.Println(x.Name, x.Age)
	}
}

func PPrint(user *User) {
	user.Name = "王五"
}

func FPrint(inter interface{}) {
	t := reflect.TypeOf(inter)
	fmt.Println(t)
	fmt.Println(t.Kind()) //获取这个接口的底层类型
	//fmt.Println(t.Elem()) //变量的原始类型

	v := reflect.ValueOf(inter)
	fmt.Println(v)
	fmt.Println(v.Type())
	fmt.Println(v.Kind()) //获取这个接口的底层类型
	//fmt.Println(v.Elem()) //变量的原始类型

	for i := 0; i < t.NumField(); i++ {
		//字段的类型，
		//字段名，
		//字段的值，
		//字段的tag
		fmt.Println(
			t.Field(i).Type,
			t.Field(i).Name,
			v.Field(i),
			t.Field(i).Tag.Get("json"))
	}

	v.FieldByName("Name").SetString("茜茜知道")
}

func main() {
	user := User{"DDDD", 21}
	Print(user)
	fmt.Println(&user)

	//PPrint(&user)
	//fmt.Println(user)
	//FPrint(user)
}
