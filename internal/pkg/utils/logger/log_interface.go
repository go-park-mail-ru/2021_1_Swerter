package logger


type LoggerModel interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
	Warn(msg string)
}