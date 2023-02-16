package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m10(c *gin.Context) {
	fmt.Println("m10 ...in")
}
func m1010(c *gin.Context) {
	fmt.Println("m1010 ...in")
}

func main() {
	router := gin.Default()

	router.Use(m10, m1010)

	router.GET("/10", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "10"})
	})

	router.GET("/11", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "11"})
	})

	router.Run(":8080")

}
