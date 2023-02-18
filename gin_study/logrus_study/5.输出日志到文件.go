package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {

	file, _ := os.OpenFile("gin_study/logrus_study/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//file, _ := os.Create("gin_study/logrus_study/info.log")

	logrus.SetOutput(io.MultiWriter(file, os.Stdout))

	logrus.Infof("你好")
	logrus.Error("你好")
	logrus.Errorf("你好 %s", "xxx")
	logrus.Errorln("你好")

}
