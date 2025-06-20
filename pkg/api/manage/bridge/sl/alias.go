package sl

type AliasRandomNewReq struct {
	Authentication string `header:"Authentication"`
	Hostname       string `form:"hostname"`
	UUID           string `form:"uuid"`
	Word           string `form:"word"`
	Note           string `json:"note"`
}

// Alias https://github.com/simple-login/app/blob/master/docs/api.md#get-apialiasesalias_id
// 实际接口返回与文档相比，需要添加部分字段
type Alias struct {
	Alias             string         `json:"alias"`         // like email
	CreationDate      string         `json:"creation_date"` // "2025-05-06 03:43:06+00:00"
	CreationTimestamp int64          `json:"creation_timestamp"`
	DisablePgp        bool           `json:"disable_pgp"`
	Email             string         `json:"email"`
	Name              string         `json:"name"`
	Enabled           bool           `json:"enabled"`
	Id                int64          `json:"id"` // 27229827
	Mailbox           MailBox        `json:"mailbox"`
	Mailboxes         []MailBox      `json:"mailboxes"` // 与 Mailbox 同时出现
	LatestActivity    LatestActivity `json:"latest_activity"`
	NbBlock           int            `json:"nb_block"`
	NbForward         int            `json:"nb_forward"`
	NbReply           int            `json:"nb_reply"`
	Note              string         `json:"note"`
	Pinned            bool           `json:"pinned"`
	SupportPgp        bool           `json:"support_pgp"`
}

type MailBox struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
}

type LatestActivity struct {
	Action    string                `json:"action"`
	Contact   LatestActivityContact `json:"contact"`
	Timestamp int                   `json:"timestamp"`
}

type LatestActivityContact struct {
	Email        string `json:"email"`
	Name         any    `json:"name"`
	ReverseAlias string `json:"reverse_alias"`
}

type ErrorResp struct {
	Error string `json:"error"`
}
