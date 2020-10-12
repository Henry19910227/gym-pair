package global

type ILogger interface {
	Trace(key string, value interface{}, msg string)
	Debug(key string, value interface{}, msg string)
	Info(key string, value interface{}, msg string)
	Warn(key string, value interface{}, msg string)
	Error(key string, value interface{}, msg string)
	Fatal(key string, value interface{}, msg string)
	Panic(key string, value interface{}, msg string)
}

var (
	Logger ILogger
)
