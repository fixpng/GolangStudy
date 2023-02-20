package main

import (
	"GO20230123/gin_study/logrus_study/gin_logrus/log"
	"GO20230123/gin_study/logrus_study/gin_logrus/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log.InitFile("gin_study/logrus_study/gin_logrus/logs", "server")
	router := gin.New()
	router.Use(middleware.LogMiddleware())

	router.GET("/", func(c *gin.Context) {
		logrus.Info("来了")
		c.JSON(200, gin.H{"msg": "你好"})
	})

	router.Run(":8081")

}
