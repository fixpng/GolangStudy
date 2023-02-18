package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
)

const (
	ccBlack  = 0
	ccRed    = 1
	ccGreen  = 2
	ccYellow = 3
	ccBlue   = 4
	ccPurple = 5
	ccCyan   = 6
	ccGray   = 7
)

type MyFormatter struct {
	Prefix     string
	TimeForMat string
}

func (f MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// 设置颜色
	var color int
	switch entry.Level {
	case logrus.ErrorLevel:
		color = ccRed
	case logrus.WarnLevel:
		color = ccYellow
	case logrus.InfoLevel:
		color = ccBlue
	case logrus.DebugLevel:
		color = ccCyan
	default:
		color = ccGray
	}

	// 设置buffer
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	// 时间格式化
	formatTime := entry.Time.Format(f.TimeForMat)
	// 文件的行号
	fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
	// 设置格式
	fmt.Fprintf(b, "[%s] \033[3%dm[%s]\033[0m [%s] [%s] %s\n", f.Prefix, color, entry.Level, formatTime, fileVal, entry.Message)

	return b.Bytes(), nil
}

func main() {

	logrus.SetReportCaller(true) //显示行号
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&MyFormatter{Prefix: "GORM", TimeForMat: "2006-01-02 15:04:05"})

	logrus.Info("你好")
	logrus.Error("你好")
	logrus.Warn("你好")
	logrus.Debug("你好")
}
