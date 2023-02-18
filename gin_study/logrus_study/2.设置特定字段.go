package main

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.WithField("app", "study").WithField("service", "logrus")

	log = log.WithFields(logrus.Fields{
		"user_id": "茜茜",
		"ip":      "192.168.200.254",
	})

	log.Errorf("你好")
}
