package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required"`
			Age  int    `json:"age" binding:"required"`
		}

		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": user})

	})
	router.Run(":8080")
}
