package logger

import (
	"github.com/spf13/viper"
)

// GPLogSetting ...
type GPLogSetting struct {
	vp *viper.Viper
}

// NewGPLogSetting ...
func NewGPLogSetting(filename string) (*GPLogSetting, error) {
	vp := viper.New()
	vp.SetConfigFile(filename)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &GPLogSetting{vp}, nil
}

// GetLogFilePath ...
func (setting *GPLogSetting) GetLogFilePath() string {
	return setting.vp.GetString("App.LogFilePath")
}

// GetLogFileName ...
func (setting *GPLogSetting) GetLogFileName() string {
	return setting.vp.GetString("App.LogFileName")
}

// GetLogFileExt ...
func (setting *GPLogSetting) GetLogFileExt() string {
	return setting.vp.GetString("App.LogFileExt")
}

// GetRunMode ...
func (setting *GPLogSetting) GetRunMode() string {
	return setting.vp.GetString("Server.RunMode")
}
