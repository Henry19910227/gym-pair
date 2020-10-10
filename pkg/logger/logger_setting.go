package logger

import (
	"github.com/spf13/viper"
)

// LoggerViperSetting ...
type LoggerViperSetting struct {
	vp *viper.Viper
}

func NewLoggerSetting(filename string) (*LoggerViperSetting, error) {
	vp := viper.New()
	vp.SetConfigFile(filename)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &LoggerViperSetting{vp}, nil
}

// GetLogFilePath ...
func (setting *LoggerViperSetting) GetLogFilePath() string {
	return setting.vp.GetString("App.LogFilePath")
}

// GetLogFileName ...
func (setting *LoggerViperSetting) GetLogFileName() string {
	return setting.vp.GetString("App.LogFileName")
}

// GetLogFileExt ...
func (setting *LoggerViperSetting) GetLogFileExt() string {
	return setting.vp.GetString("App.LogFileExt")
}
