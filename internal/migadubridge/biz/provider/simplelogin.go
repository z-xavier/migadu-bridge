package provider

import "context"

type SimpleLogin struct{}

func NewSimpleLogin() *SimpleLogin {
	return &SimpleLogin{}
}

func (sl *SimpleLogin) AliasRandomNew(ctx context.Context, domain, desc string) error {
	return nil
}
