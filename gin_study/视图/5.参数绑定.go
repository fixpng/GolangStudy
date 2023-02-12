package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	router := gin.Default()

	// JSON
	router.POST("/", func(c *gin.Context) {

		var userinfo UserInfo
		err := c.ShouldBindJSON(&userinfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(0, gin.H{"msg": "你出错了"})
			return
		}
		c.JSON(0, userinfo)
	})
	// tag 对应为query
	router.POST("/query", func(c *gin.Context) {

		var userinfo UserInfo
		err := c.ShouldBindQuery(&userinfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(0, gin.H{"msg": "你出错了"})
			return
		}
		c.JSON(0, userinfo)
	})
	// tag 对应为uri
	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {

		var userinfo UserInfo
		err := c.ShouldBindUri(&userinfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(0, gin.H{"msg": "你出错了"})
			return
		}
		c.JSON(0, userinfo)
	})
	//绑定form ，默认的tag就是form
	router.POST("/form", func(c *gin.Context) {

		var userinfo UserInfo
		err := c.ShouldBind(&userinfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(0, gin.H{"msg": "你出错了"})
			return
		}
		c.JSON(0, userinfo)
	})

	router.Run(":8080")

}
