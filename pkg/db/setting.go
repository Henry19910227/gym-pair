package db

import (
	"github.com/spf13/viper"
)

// MysqlViperSetting ...
type MysqlViperSetting struct {
	vp *viper.Viper
}

// NewMysqlSetting ...
func NewMysqlSetting(filename string) (*MysqlViperSetting, error) {
	vp := viper.New()
	vp.SetConfigFile(filename)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &MysqlViperSetting{vp}, nil
}

// GetUserName ...
func (setting *MysqlViperSetting) GetUserName() string {
	return setting.vp.GetString("Database.UserName")
}

// GetPassword ...
func (setting *MysqlViperSetting) GetPassword() string {
	return setting.vp.GetString("Database.Password")
}

// GetHost ...
func (setting *MysqlViperSetting) GetHost() string {
	return setting.vp.GetString("Database.Host")
}

// GetDatabase ...
func (setting *MysqlViperSetting) GetDatabase() string {
	return setting.vp.GetString("Database.DBName")
}

// SetHost ...
func (setting *MysqlViperSetting) SetHost(host string) {
	setting.vp.Set("Database.Host", host)
}
