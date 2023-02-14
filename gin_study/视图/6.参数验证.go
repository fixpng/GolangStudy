package main

import (
	"github.com/gin-gonic/gin"
)

//bind绑定器 gin内置验证器
//http://docs.fengfengzhidao.com/#/docs/Gin%E6%A1%86%E6%9E%B6%E6%96%87%E6%A1%A3/4.bind%E7%BB%91%E5%AE%9A%E5%99%A8

type SignUserInfo struct {
	Name       string   `json:"name" binding:"required,min=4,max=6"` //bind绑定 binding:"required" 不能为空，且不能没有这个字段
	Age        int      `json:"age" binding:"lt=30,gt=18"`
	Password   string   `json:"password"`
	RePassword string   `json:"re_password" binding:"eqfield=Password"` //等于其他字段的值
	Sex        string   `json:"sex"  binding:"oneof=man woman"`         //枚举
	LikeList   []string `json:"like_list" binding:"dive,required,min=1,startswith=like"`
	IP         string   `json:"ip" binding:"ip"`
	Url        string   `json:"url" binding:"url"`
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {

		var user SignUserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": user})
	})

	router.Run(":8080")
}
