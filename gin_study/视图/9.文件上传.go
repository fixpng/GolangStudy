package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

var dstPath = "./gin_study/uploads/"

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// 单文件上传 1
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := dstPath + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst) // 文件对象  文件路径，注意要从项目根路径开始写

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	// 服务端保存文件
	router.POST("/uploadSave", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// 读取文件中的数据，返回文件对象
		fileRead, _ := file.Open()
		dst := dstPath + file.Filename
		// 创建一个文件
		out, err := os.Create(dst)
		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()
		// 拷贝文件对象到out中
		io.Copy(out, fileRead)
	})

	// 多文件上传
	router.POST("/uploads", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"] // 注意这里名字不要对不上了

		for _, file := range files {
			log.Println(file.Filename)
			// 上传文件至指定目录
			c.SaveUploadedFile(file, dstPath+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run(":8080")
}
