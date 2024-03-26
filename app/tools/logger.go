package tools

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger *logrus.Logger

func NewLogger() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)

	// 同时写到多个输出
	w1 := os.Stdout
	w2, _ := os.OpenFile("./vote.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	Logger.SetOutput(io.MultiWriter(w1, w2))
	Logger.SetFormatter(&logrus.JSONFormatter{})
}
