package migadubridge

import (
	"github.com/bytedance/sonic"
	"gorm.io/gorm"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/config"
	"migadu-bridge/internal/pkg/db"
	"migadu-bridge/internal/pkg/model"
)

// initStore 读取 db 配置，创建 gorm.DB 实例，并初始化 migadu store 层.
func initStore() error {
	var (
		ins *gorm.DB
	)

	dbConfigString, err := sonic.MarshalString(config.GetConfig().DB)
	if err != nil {
		return err
	}

	if config.GetConfig().DB.Driver == "" {
		config.GetConfig().DB.Driver = "sqlite"
	}

	switch config.GetConfig().DB.Driver {
	case "sqlite":
		var dbOptions db.SqliteOptions
		if err = sonic.UnmarshalString(dbConfigString, &dbOptions); err != nil {
			return err
		}
		ins, err = db.NewSqlite(&dbOptions)
		if err != nil {
			return err
		}
	case "mysql":
		var dbOptions db.MySQLOptions
		if err = sonic.UnmarshalString(dbConfigString, &dbOptions); err != nil {
			return err
		}
		ins, err = db.NewMySQL(&dbOptions)
		if err != nil {
			return err
		}
	}

	_ = store.NewStore(ins)

	err = store.S.DB().AutoMigrate(
		&model.CallLog{},
		&model.Token{},
	)
	if err != nil {
		return err
	}

	return nil
}
