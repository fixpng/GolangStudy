package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type MyHook struct {
}

func (hook MyHook) Levels() []logrus.Level {
	//return logrus.AllLevels
	return []logrus.Level{logrus.ErrorLevel}
}
func (hook MyHook) Fire(entry *logrus.Entry) error {
	//entry.Data["app"] = "shabi"
	file, _ := os.OpenFile("gin_study/logrus_study/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	line, _ := entry.String()
	file.Write([]byte(line))
	return nil
}

func main() {
	logrus.AddHook(&MyHook{})

	logrus.Warnln("你好")
	logrus.Errorf("你好")
}
