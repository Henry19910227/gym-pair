package setting

import (
	"github.com/spf13/viper"
)

// ViperSetting ...
type ViperSetting struct {
	vp *viper.Viper
}

// NewViperSetting ...
func NewViperSetting() (*ViperSetting, error) {
	vp := viper.New()
	vp.SetConfigFile("./config/config.yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &ViperSetting{vp}, nil
}

// GetUserName ...
func (setting *ViperSetting) GetUserName() string {
	return setting.vp.GetString("Database.UserName")
}

// GetPassword ...
func (setting *ViperSetting) GetPassword() string {
	return setting.vp.GetString("Database.Password")
}

// GetHost ...
func (setting *ViperSetting) GetHost() string {
	return setting.vp.GetString("Database.Host")
}

// GetDatabase ...
func (setting *ViperSetting) GetDatabase() string {
	return setting.vp.GetString("Database.DBName")
}
