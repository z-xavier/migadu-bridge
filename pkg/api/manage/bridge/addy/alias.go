package addy

type AliasFormat string

const (
	AliasFormatRandomCharacters AliasFormat = "random_characters"
	AliasFormatUUID             AliasFormat = "uuid"
	AliasFormatRandomWords      AliasFormat = "random_words"
	AliasFormatCustom           AliasFormat = "custom"
)

// Alias https://app.addy.io/docs/#aliases-POSTapi-v1-aliases
type Alias struct {
	Id              string `json:"id"`
	UserId          string `json:"user_id"`
	AliasableId     any    `json:"aliasable_id"`
	AliasableType   any    `json:"aliasable_type"`
	LocalPart       string `json:"local_part"`
	Extension       any    `json:"extension"`
	Domain          string `json:"domain"`
	Email           string `json:"email"`
	Active          bool   `json:"active"`
	Description     string `json:"description"`
	FromName        any    `json:"from_name"`
	EmailsForwarded int    `json:"emails_forwarded"`
	EmailsBlocked   int    `json:"emails_blocked"`
	EmailsReplied   int    `json:"emails_replied"`
	EmailsSent      int    `json:"emails_sent"`
	Recipients      []any  `json:"recipients"`
	LastForwarded   string `json:"last_forwarded"`
	LastBlocked     any    `json:"last_blocked"`
	LastReplied     any    `json:"last_replied"`
	LastSent        any    `json:"last_sent"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}

type CreateAliasReq struct {
	Authorization  string      `header:"Authorization"`
	XRequestedWith string      `header:"X-Requested-With"`
	Domain         string      `json:"domain"`
	Description    string      `json:"description"`
	Format         AliasFormat `json:"format"`
	LocalPart      string      `json:"local_part"`
	RecipientIds   []string    `json:"recipient_ids"`
}

type CreateAliasResp struct {
	Data *Alias `json:"data"`
}
