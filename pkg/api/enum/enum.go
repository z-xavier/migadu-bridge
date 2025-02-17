package enum

type TokenStatus uint8

const (
	TokenStatusInactive TokenStatus = iota + 1
	TokenStatusActive
	TokenStatusPause
)

type ProviderEnum string

const (
	ProviderEnumAddy        ProviderEnum = "addy"
	ProviderEnumSimpleLogin ProviderEnum = "sl"
)
