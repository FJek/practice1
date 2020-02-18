package config

import "fmt"

// 对外实例
var MySQL MySQLConfig

type MySQLConfig struct {

	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     int    `env:"DB_PORT" envDefault:"3306"`
	Username string `env:"DB_USERNAME" envDefault:"root"`
	Password string `env:"DB_PASSWORD" envDefault:"123456"`
	Database string `env:"DB_DATABASE" envDefault:"test"`
	// EnableLog 是否启用Gorm SQL执行log
	EnableLog bool `env:"SQL_LOG_ENABLE" envDefault:"false"`
	// DBPoolMax 连接池最大数量
	DBPoolMax int `env:"DB_POOL_MAX" envDefault:"50"`
	// DBPoolMaxIdle 连接池最大闲置数量
	DBPoolMaxIdle int `env:"DB_POOL_MAX_IDLE" envDefault:"20"`
}

// GetDsn 获取MySQL连接DSN
func (config *MySQLConfig) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
}