package model

import (
	"time"

	"migadu-bridge/pkg/api/enum"
)

type Token struct {
	Model
	TargetEmail  string            `json:"target_email,omitempty"`
	MockProvider enum.ProviderEnum `json:"mock_provider,omitempty"`
	Description  string            `json:"description,omitempty"`
	Token        string            `json:"token,omitempty"`
	ExpiryAt     time.Time         `json:"expiry_at"`
	LastCalledAt time.Time         `json:"last_called_at"`
	Status       enum.TokenStatus  `json:"status,omitempty"`
	CallLogs     []CallLog         `gorm:"foreignKey:TokenId"`
}
