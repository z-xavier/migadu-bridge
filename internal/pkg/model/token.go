package model

import (
	"time"

	"gorm.io/gorm"

	"migadu-bridge/pkg/api/enum"
)

type Token struct {
	Id           string
	TargetEmail  string
	MockProvider enum.ProviderEnum
	Description  string
	Token        string
	ExpiryAt     time.Time
	LastCalledAt time.Time
	Status       enum.TokenStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
