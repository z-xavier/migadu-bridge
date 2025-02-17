package call_log

import "migadu-bridge/internal/migadubridge/store"

type CallLogBiz interface{}

type callLogBiz struct {
	ds store.IStore
}

func New(ds store.IStore) CallLogBiz {
	return &callLogBiz{ds: ds}
}
