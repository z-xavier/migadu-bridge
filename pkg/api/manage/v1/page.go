package v1

type PageReqHeader struct {
	Page     uint64 `form:"page" json:"page" binding:"omitempty,min=1"`
	PageSize uint64 `form:"pageSize" json:"pageSize" binding:"omitempty,min=1,max=200"`
}

type Page struct {
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"pageSize"`
	Total    uint64 `json:"total"`
}
