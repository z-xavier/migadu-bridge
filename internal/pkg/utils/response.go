package utils

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

const (
	initBufCap = 4 * 1024
	maxBufCap  = 512 * 1024
)

var respBodyBuf = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, initBufCap))
	}}

func NewRespBuf() *bytes.Buffer {
	buf := respBodyBuf.Get().(*bytes.Buffer)
	return buf
}

func ReleaseRespBuf(buf *bytes.Buffer) {
	if buf.Cap() > maxBufCap {
		return
	}
	buf.Reset()
	respBodyBuf.Put(buf)
}

func UnmarshalFromResponse[T any](r *http.Response) (*T, error) {
	buf := NewRespBuf()
	defer ReleaseRespBuf(buf)

	if _, err := buf.ReadFrom(r.Body); err != nil {
		return nil, errors.WithMessagef(err, "failed to copy body")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Errorf("StatusCode != 200, data: %s, StatusCode: %d", buf.String(), r.StatusCode)
	}
	v := new(T)
	if err := sonic.ConfigStd.Unmarshal(buf.Bytes(), v); err != nil {
		return nil, errors.WithMessagef(err, "UnmarshalFromResponse, data: %s", buf.String())
	}
	return v, nil
}
