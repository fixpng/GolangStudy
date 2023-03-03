package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) SetAge(age int) {
	u.Age = age
}

func (u *User) SetName(name string) {
	(*u).Name = name
}

func main() {
	user := User{"张三", 56}

	user.SetName("张思")
	user.SetAge(20)

	fmt.Println(user)
}
