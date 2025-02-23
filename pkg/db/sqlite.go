package db

import (
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SqliteOptions 定义 Sqlite 数据库的选项.
type SqliteOptions struct {
	Path     string
	WAL      bool
	LogLevel int
}

// DSN 从 SqliteOptions 返回 DSN.
func (o *SqliteOptions) DSN() string {
	return path.Join(o.Path, "sqlite.db")
}

// NewSqlite 使用给定的选项创建一个新的 gorm 数据库实例.
func NewSqlite(opts *SqliteOptions) (*gorm.DB, error) {
	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}

	dsn := opts.DSN()
	if dsn == "" {
		dsn = "sqlite.db"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logLevel),
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 开启 WAL 模式
	if opts.WAL {
		_, err = sqlDB.Exec("PRAGMA journal_mode=WAL;")
	}

	return db, nil
}
