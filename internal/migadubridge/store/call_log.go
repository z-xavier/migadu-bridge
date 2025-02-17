package store

import (
	"context"

	"gorm.io/gorm"

	"migadu-bridge/internal/pkg/model"
)

type CallLogStore interface {
	Create(ctx context.Context, callLog *model.CallLog) (string, error)
}

type callLogStore struct {
	db *gorm.DB
}

func NewCallLogStore(db *gorm.DB) CallLogStore {
	return &callLogStore{db}
}

func (t *callLogStore) Create(ctx context.Context, callLog *model.CallLog) (string, error) {
	return "", nil
}
