package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Error("出错了")
	logrus.Warningln("警告")
	logrus.Infof("信息")
	logrus.Debugf("debug")

	logrus.Println("打印")

	fmt.Println(logrus.GetLevel())

}
