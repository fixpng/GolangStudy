package main

import "fmt"

func main() {
	type User struct {
		Name     string
		Age      int
		password string
	}

	type Account struct {
		User
		money float32
	}

	var ac Account = Account{
		money: 21.9,
		User: User{
			Name:     "王五",
			Age:      5,
			password: "#$FG%3r",
		},
	}
	fmt.Printf("ac %#v,%T\n", ac, ac)
}
