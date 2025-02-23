package v1

import "migadu-bridge/pkg/api/enum"

type CallLog struct {
	Id           string            `json:"id"`
	TokenId      string            `json:"tokenId"`
	TargetEmail  string            `json:"targetEmail"`
	MockProvider enum.ProviderEnum `json:"mockProvider"`
	GenAlias     string            `json:"genAlias"`
	RequestPath  string            `json:"requestPath"`
	RequestIp    string            `json:"requestIp"`
	RequestAt    int64             `json:"requestAt"`
}

type ListCallLogReq struct {
	PageReqHeader
	TargetEmail    string            `form:"targetEmail" json:"targetEmail"`
	MockProvider   enum.ProviderEnum `form:"mockProvider" json:"mockProvider"`
	RequestPath    string            `form:"requestPath" json:"requestPath"`
	RequestIp      string            `form:"requestIp" json:"requestIp"`
	RequestAtBegin int64             `form:"requestAtBegin" json:"requestAtBegin"`
	RequestAtEnd   int64             `form:"requestAtEnd" json:"requestAtEnd"`
}

type ListCallLogResp struct {
	Page
	List []*CallLog `json:"list"`
}
