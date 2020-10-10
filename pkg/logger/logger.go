package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LoggerSetting interface {
	GetLogFilePath() string
	GetLogFileName() string
	GetLogFileExt() string
}

type Logger struct {
	loggerItem *logrus.Logger
}

// NewLogger ...
func NewLogger(setting LoggerSetting) (*Logger, error) {
	// file, err := os.OpenFile("/Users/henry/go/src/github.com/Henry19910227/gym-pair/storage/app.log", os.O_WRONLY|os.O_CREATE, 0777)
	file, err := os.OpenFile(setting.GetLogFilePath()+setting.GetLogFileName()+"."+setting.GetLogFileExt(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	loggerItem := logrus.New()
	loggerItem.SetOutput(file)

	return &Logger{loggerItem}, nil
}

func (logger *Logger) Info(msg string) {
	logger.loggerItem.WithField("GetAll", "All").Info(msg)
}
