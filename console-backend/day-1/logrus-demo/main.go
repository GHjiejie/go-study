package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// 创建一个新的日志记录器
	log := logrus.New()

	// 设置日志级别
	log.SetLevel(logrus.DebugLevel)

	// 记录一些日志
	log.Debug("This is a debug log")
	log.Info("This is an info log")
	log.Warn("This is a warning log")
	log.Error("This is an error log")
	log.Fatal("This is a fatal log")
}
