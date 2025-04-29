package sl

type AliasRandomNewReq struct {
	Authentication string `header:"Authentication"`
	Hostname       string `form:"hostname"`
	UUID           string `form:"uuid"`
	Word           string `form:"word"`
	Note           string `json:"note"`
}

// Alias https://github.com/simple-login/app/blob/master/docs/api.md#get-apialiasesalias_id
type Alias struct {
	CreationDate      string         `json:"creation_date"`
	CreationTimestamp int64          `json:"creation_timestamp"`
	Email             string         `json:"email"`
	Name              string         `json:"name"`
	Enabled           bool           `json:"enabled"`
	Id                int64          `json:"id"`
	Mailbox           MailBox        `json:"mailbox"`
	Mailboxes         []MailBox      `json:"mailboxes"`
	LatestActivity    LatestActivity `json:"latest_activity"`
	NbBlock           int            `json:"nb_block"`
	NbForward         int            `json:"nb_forward"`
	NbReply           int            `json:"nb_reply"`
	Note              any            `json:"note"`
	Pinned            bool           `json:"pinned"`
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
