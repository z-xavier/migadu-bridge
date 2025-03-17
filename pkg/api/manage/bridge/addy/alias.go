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
	Id              string        `json:"id"`
	UserId          string        `json:"user_id"`
	AliasableId     interface{}   `json:"aliasable_id"`
	AliasableType   interface{}   `json:"aliasable_type"`
	LocalPart       string        `json:"local_part"`
	Extension       interface{}   `json:"extension"`
	Domain          string        `json:"domain"`
	Email           string        `json:"email"`
	Active          bool          `json:"active"`
	Description     string        `json:"description"`
	FromName        interface{}   `json:"from_name"`
	EmailsForwarded int           `json:"emails_forwarded"`
	EmailsBlocked   int           `json:"emails_blocked"`
	EmailsReplied   int           `json:"emails_replied"`
	EmailsSent      int           `json:"emails_sent"`
	Recipients      []interface{} `json:"recipients"`
	LastForwarded   string        `json:"last_forwarded"`
	LastBlocked     interface{}   `json:"last_blocked"`
	LastReplied     interface{}   `json:"last_replied"`
	LastSent        interface{}   `json:"last_sent"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
	DeletedAt       interface{}   `json:"deleted_at"`
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
