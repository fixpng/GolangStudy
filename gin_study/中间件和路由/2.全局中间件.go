package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func m10(c *gin.Context) {
	fmt.Println("m10 ...in")
	c.Set("name", "茜茜")
	c.Set("user", User{
		Name: "XX",
		Age:  20,
	})

}
func m1010(c *gin.Context) {
	fmt.Println("m1010 ...in")
}

func main() {
	router := gin.Default()

	router.Use(m10, m1010)

	router.GET("/10", func(c *gin.Context) {
		name, _ := c.Get("name")
		fmt.Println(name)
		_user, _ := c.Get("user")
		user := _user.(User)
		fmt.Println(user.Age, user.Name)

		c.JSON(200, gin.H{"msg": "10"})
	})

	router.GET("/11", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "11"})
	})

	router.Run(":8080")

}
