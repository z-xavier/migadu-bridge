package bridge

import (
	"context"

	"migadu-bridge/internal/pkg/rwords"
)

type AddyBridge struct {
}

func (ab *AddyBridge) AliasRandomNew(ctx context.Context, domain, desc string) error {
	alias, err := rwords.GetRWords(false, true)
	if err != nil {
		return err
	}

}
