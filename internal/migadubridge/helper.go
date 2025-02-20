package migadubridge

import (
	"github.com/spf13/viper"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/pkg/db"
)

// initStore 读取 db 配置，创建 gorm.DB 实例，并初始化 migadu store 层.
func initStore() error {
	dbOptions := &db.MySQLOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}

	ins, err := db.NewMySQL(dbOptions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)

	return nil
}
