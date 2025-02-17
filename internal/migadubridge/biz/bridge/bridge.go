package bridge

import "migadu-bridge/internal/migadubridge/store"

type BridgeBiz interface{}

type bridgeBiz struct {
	ds store.IStore
}

func New(ds store.IStore) BridgeBiz {
	return &bridgeBiz{ds: ds}
}
