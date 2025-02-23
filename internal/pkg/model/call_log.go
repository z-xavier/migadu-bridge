package model

import "time"

type CallLog struct {
	Model
	TokenId     string    `json:"token_id"`
	RequestPath string    `json:"request_path"`
	GenAlias    string    `json:"gen_alias"`
	RequestAt   time.Time `json:"request_at"`
	RequestIp   string    `json:"request_ip"`
}
