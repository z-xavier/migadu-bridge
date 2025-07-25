package db

import (
	"path"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"migadu-bridge/internal/pkg/log"
)

// SqliteOptions 定义 Sqlite 数据库的选项.
type SqliteOptions struct {
	Path     string `json:"path"`
	WAL      bool   `json:"wal"`
	LogLevel int32  `json:"log-level"`
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

	log.WithField("sqlite", dsn).Infow("Open sqlite database")

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
