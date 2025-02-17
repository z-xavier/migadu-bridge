package v1

type Log struct {
	Id           string `json:"id"`
	TokenId      string `json:"tokenId"`
	TargetEmail  string `json:"targetEmail"`
	MockProvider string `json:"mockProvider"`
	Path         string `json:"path"`
	Token        string `json:"token"`
	Ip           string `json:"ip"`
	CallTime     uint64 `json:"callTime"`
}

type ListLogResp struct {
	Page
	List []Log `json:"list"`
}
