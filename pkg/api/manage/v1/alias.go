package v1

import "migadu-bridge/pkg/api/enum"

type Alias struct {
	Id           int64             `json:"id"`
	TargetEmail  string            `json:"targetEmail"`
	Alias        string            `json:"alias"`
	CallLogId    string            `json:"callLogId"`
	TokenId      string            `json:"tokenId"`
	MockProvider enum.ProviderEnum `json:"mockProvider"`
}

type ListAliasReq struct {
	PageReqHeader
}

type ListAliasResp struct {
	Page
	List []*Alias `json:"list"`
}
