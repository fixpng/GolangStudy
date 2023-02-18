package main

import "github.com/sirupsen/logrus"

func main() {

	//默认是 TextFormatter{} 可改 JSONFormatter{}
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetLevel(logrus.DebugLevel)

	logrus.Errorf("你好")
	logrus.Info("你好")
	logrus.Warnln("你好")
	logrus.Debugf("你好")
	logrus.Println("你好")
}
