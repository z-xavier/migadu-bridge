package v1

import (
	"migadu-bridge/pkg/api/enum"
)

type Token struct {
	Id           string            `json:"id"`
	TargetEmail  string            `json:"targetEmail"`
	MockProvider enum.ProviderEnum `json:"mockProvider"`
	Description  string            `json:"description"`
	Token        string            `json:"token"`
	ExpiryAt     int64             `json:"expiryAt"`
	LastCalledAt int64             `json:"lastCalledAt"`
	Status       enum.TokenStatus  `json:"status"`
	CreatedAt    int64             `json:"createdAt"`
	UpdatedAt    int64             `json:"updatedAt"`
}

type CreateTokenReq struct {
	TargetEmail  string            `json:"targetEmail" binding:"required,email"`
	MockProvider enum.ProviderEnum `json:"mockProvider" binding:"required"`
	Description  string            `json:"description" binding:"max=1024"`
	ExpiryAt     int64             `json:"expiryAt" binding:"required,min=0"`
}

type CreateTokenResp struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

type UpdateTokenReq struct {
	Description string           `json:"description"`
	ExpiryAt    int64            `json:"expiryAt"`
	Status      enum.TokenStatus `json:"status"`
}

type ListTokenReq struct {
	PageReqHeader
	Id                string            `form:"id" json:"id"`
	TargetEmail       string            `form:"targetEmail" json:"targetEmail"`
	MockProvider      enum.ProviderEnum `form:"mockProvider" json:"mockProvider"`
	Description       string            `form:"description" json:"description"`
	ExpiryAtBegin     int64             `form:"expiryAtBegin" json:"expiryAtBegin"`
	ExpiryAtEnd       int64             `form:"expiryAtEnd" json:"expiryAtEnd"`
	LastCalledAtBegin int64             `form:"lastCalledAtBegin" json:"lastCalledAtBegin"`
	LastCalledAtEnd   int64             `form:"lastCalledAtEnd" json:"lastCalledAtEnd"`
	Status            enum.TokenStatus  `form:"status" json:"status"`
}

type GetTokenResp struct {
	Token
}

type ListTokenResp struct {
	Page
	List []*Token `json:"list"`
}
