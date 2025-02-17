package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	// 全局变量，方便其它包直接调用已初始化好的 S 实例.
	S *datastore
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
	DB() *gorm.DB
	//Users() UserStore
	//Posts() PostStore
}

// datastore 是 IStore 的一个具体实现.
type datastore struct {
	db *gorm.DB
}

// NewStore 创建一个 IStore 类型的实例.
func NewStore(db *gorm.DB) *datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{db}
	})
	return S
}

func (d datastore) DB() *gorm.DB {
	return d.db
}
