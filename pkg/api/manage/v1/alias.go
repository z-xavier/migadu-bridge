package v1

import "migadu-bridge/pkg/api/enum"

type Alias struct {
	Id               int64             `json:"id"`
	TargetEmail      string            `json:"targetEmail"`
	Alias            string            `json:"alias"` // 别名全称
	CallLogId        string            `json:"callLogId"`
	TokenId          string            `json:"tokenId"`
	MockProvider     enum.ProviderEnum `json:"mockProvider"`
	Description      string            `json:"description"`
	Expireable       bool              `json:"expireable"`
	ExpiresOn        string            `json:"expires_on"` // 过期时间
	IsInternal       bool              `json:"is_internal"`
	RemoveUponExpiry bool              `json:"remove_upon_expiry"`
	RequestAt        int64             `json:"requestAt"`
}

type ListAliasReq struct {
	PageReqHeader
}

type ListAliasResp struct {
	Page
	List []*Alias `json:"list"`
}
