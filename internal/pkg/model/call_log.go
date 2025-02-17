package model

import "time"

type CallLog struct {
	Id          string    `json:"id"`
	TokenId     uint64    `json:"token_id"`
	RequestPath string    `json:"request_path"`
	GenAlias    string    `json:"gen_alias"`
	RequestTime time.Time `json:"request_time"`
	RequestIp   string    `json:"request_ip"`
	UpdatedTime time.Time `json:"updated_time"`
	CreatedTime time.Time `json:"created_time"`
}
