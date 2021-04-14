package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Logger struct {
	Log *logrus.Logger
}

func NewLogger() *Logger {
	newLogger := SetLogger()
	return &Logger{
		Log: newLogger,
	}
}

func (l *Logger) Info(msg string) {
	l.Log.Info(msg)
}

func (l *Logger) Debug(msg string) {
	l.Log.Debug(msg)
}

func (l *Logger) Error(msg string) {
	l.Log.Error(msg)
}

func (l *Logger) Warn(msg string) {
	l.Log.Warn(msg)
}

type myFormatter struct {
	logrus.TextFormatter
}

func SetLogger() *logrus.Logger {
	//f, _ := os.OpenFile("logrus.txt", os.O_CREATE|os.O_WRONLY, 0666)
	logger := &logrus.Logger{
		Out:   io.MultiWriter(os.Stderr), //можно добавтиь файл
		Level: logrus.DebugLevel,
		Formatter: &myFormatter{
			logrus.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "2006-01-02 15:04:05",
				ForceColors:            true,
				DisableLevelTruncation: true,
			},
		},
	}
	return logger
}
