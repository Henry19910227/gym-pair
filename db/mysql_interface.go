package db

// MysqlSetting ...
type MysqlSetting interface {
	GetUserName() string
	GetPassword() string
	GetHost() string
	GetDatabase() string
}
