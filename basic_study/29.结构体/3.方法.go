package main

//func (u User) PrintName() string {
//	fmt.Println("PrintName方法： \n", u)
//	fmt.Printf("PrintName方法内部: %p\n", &u)
//	return u.Name
//}

func main() {

	type User struct {
		Name     string
		Age      int
		password string
	}

	// Account 继承也会把方法继承
	type Account struct {
		Money float32
	}

	//user := User{"王二麻子", 23, "sdg#$%d12"}
	//fmt.Printf("main: %p\n", &user)
	//name := user.PrintName() // 是值传递
	//fmt.Println(name)
}
