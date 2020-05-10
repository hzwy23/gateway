package logger

import (
	"github.com/sirupsen/logrus"
	"log"
)

var lg = logrus.New()

func Debug(args ...interface{}) {
	lg.Debug(args...)
}

func Info(args ...interface{}) {
	lg.Info(args...)
}

func Error(args ...interface{}) {
	lg.Error(args...)
}

func Trace(args ...interface{}) {
	lg.Trace(args...)
}

func Fatal(args ...interface{}) {
	lg.Fatal(args...)
}

func init() {

	lg.Formatter = &logrus.JSONFormatter{}
	log.SetOutput(lg.Writer())
}
