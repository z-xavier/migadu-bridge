package v1

import (
	"migadu-bridge/pkg/api/enum"
)

type Token struct {
	Id           uint64            `json:"id"`
	TargetEmail  string            `json:"targetEmail"`
	MockProvider enum.ProviderEnum `json:"mockProvider"`
	Description  string            `json:"description"`
	Token        string            `json:"token"`
	ExpiryAt     uint64            `json:"expiryAt"`
	LastCalledAt uint64            `json:"lastCalledAt"`
	Status       enum.TokenStatus  `json:"status"`
	CreatedAt    uint64            `json:"createdAt"`
	UpdatedAt    uint64            `json:"updatedAt"`
}

type CreateTokenReq struct {
	TargetEmail  string            `json:"targetEmail" validate:"required,email"`
	MockProvider enum.ProviderEnum `json:"mockProvider" validate:"required"`
	Description  string            `json:"description" validate:"max=1024"`
	ExpiryAt     uint64            `json:"expiryAt" validate:"required,datetime"`
}

type CreateTokenResp struct {
	Id uint64 `json:"id"`
}

type ListTokenResp struct {
	Page
	List []Token `json:"list"`
}
