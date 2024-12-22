package bridge

import (
	"context"

	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/internal/pkg/rwords"
)

type AddyBridge struct {
}

func (ab *AddyBridge) AliasRandomNew(ctx context.Context, domain, desc string) error {
	log.C(ctx).Infow("adding new alias", "domain", domain, "description", desc)
	client, err := MigaduClient()
	if err != nil {
		log.C(ctx).Infow("error creating client", "error", err)
		return err
	}
	// TODO
	var destinations = []string{
		"me@zxavier.com",
	}
	var localPart string
	for {
		localPart, err = rwords.GetRWords(false, true)
		if err != nil {
			return err
		}
		// TODO 404 ERROR
		alias, _ := client.GetAlias(ctx, localPart)
		if alias == nil {
			break
		}
	}
	log.C(ctx).Infow("begin adding new alias", "alias", localPart, "description", destinations)
	_, err = client.NewAlias(ctx, localPart, destinations)
	if err != nil {
		log.C(ctx).Infow("adding new alias failed", "alias", localPart, "description", destinations)
		return err
	}
	log.C(ctx).Infow("adding new alias success", "alias", localPart, "description", destinations)
	return nil
}
