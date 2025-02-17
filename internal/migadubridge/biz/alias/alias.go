package alias

import "migadu-bridge/internal/migadubridge/store"

type AliasBiz interface {
}

type aliasBiz struct {
	ds store.IStore
}

func New(ds store.IStore) AliasBiz {
	return &aliasBiz{ds: ds}
}
