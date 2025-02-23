package v1

import "migadu-bridge/pkg/api/enum"

type CallLog struct {
	Id           string            `json:"id"`
	TokenId      string            `json:"tokenId"`
	TargetEmail  string            `json:"targetEmail"`
	MockProvider enum.ProviderEnum `json:"mockProvider"`
	RequestPath  string            `json:"requestPath"`
	GenAlias     string            `json:"genAlias"`
	RequestIp    string            `json:"requestIp"`
	RequestAt    uint64            `json:"requestAt"`
}

type ListCallLogResp struct {
	Page
	List []*CallLog `json:"list"`
}
