package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

//TODO реализовать интерфейс структуру логера и интерфейс

func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		FullTimestamp: true,
	}

	return l
}
