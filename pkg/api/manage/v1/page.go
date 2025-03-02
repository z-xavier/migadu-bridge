package v1

type PageReqHeader struct {
	Page     int64    `form:"page,default=1" json:"page" binding:"omitempty,min=1"`
	PageSize int64    `form:"pageSize,default=10" json:"pageSize" binding:"omitempty,min=1,max=200"`
	OrderBy  []string `form:"orderBy" json:"orderBy"`
}

type Page struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	Total    int64 `json:"total"`
}
