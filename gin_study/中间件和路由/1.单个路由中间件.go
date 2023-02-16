package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("m1 ...")
	c.JSON(200, gin.H{"msg": "m1的响应"})
	//c.Abort() //拦截
	c.Next()
	fmt.Println("m1 ...out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 ...")
	c.JSON(200, gin.H{"msg": "m2的响应"})
	c.Next()
	fmt.Println("m2 ...out")
}

func main() {
	router := gin.Default()
	router.GET("/", m1, func(c *gin.Context) {
		fmt.Println("index")
		c.JSON(200, gin.H{"msg": "index"})
		fmt.Println("index ...out")
	}, m2)

	router.Run(":8080")

}
