package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// LoggerSetting ...
type LoggerSetting interface {
	GetLogFilePath() string
	GetLogFileName() string
	GetLogFileExt() string
	GetRunMode() string
}

// Logger ...
type Logger struct {
	RunMode string
	print   *logrus.Logger
	write   *logrus.Logger
}

// NewLogger ...
func NewLogger(setting LoggerSetting) (*Logger, error) {
	writeLog, err := newWriteLogger(setting)
	if err != nil {
		return nil, err
	}
	printLog := newPrintLogger()
	runMode := setting.GetRunMode()

	return &Logger{runMode, printLog, writeLog}, nil
}

func newPrintLogger() *logrus.Logger {
	printLog := logrus.New()
	printLog.SetFormatter(&logrus.JSONFormatter{})
	printLog.SetOutput(os.Stdout)
	return printLog
}

func newWriteLogger(setting LoggerSetting) (*logrus.Logger, error) {
	file, err := os.OpenFile(setting.GetLogFilePath()+"/"+setting.GetLogFileName()+"."+setting.GetLogFileExt(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	writeLog := logrus.New()
	writeLog.SetFormatter(&logrus.JSONFormatter{})
	writeLog.SetOutput(file)
	return writeLog, nil
}

// Trace ...
func (logger *Logger) Trace(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Trace(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Trace(msg)
	}

}

// Debug ...
func (logger *Logger) Debug(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Debug(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Debug(msg)
	}
}

// Info ...
func (logger *Logger) Info(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Info(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Info(msg)
	}
}

// Warn ...
func (logger *Logger) Warn(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Warn(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Warn(msg)
	}
}

// Error ...
func (logger *Logger) Error(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Error(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Error(msg)
	}
}

// Fatal ...
func (logger *Logger) Fatal(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Fatal(msg)
	if logger.RunMode == "release" {
		logger.write.WithField(key, value).Fatal(msg)
	}
}

// Panic ...
func (logger *Logger) Panic(key string, value interface{}, msg string) {
	logger.print.WithField(key, value).Panic(msg)
	if logger.RunMode != "debug" {
		logger.write.WithField(key, value).Panic(msg)
	}
}
