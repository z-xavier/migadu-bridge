package model

import "time"

type CallLog struct {
	Model
	TokenId     string    `json:"token_id"`
	Description string    `json:"description"`
	GenAlias    string    `json:"gen_alias"`
	RequestPath string    `json:"request_path"`
	RequestIp   string    `json:"request_ip"`
	RequestAt   time.Time `json:"request_at"`
}
