package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"

	"migadu-bridge/internal/pkg/log"
)

type ContentType string

const (
	ContentTypeHeader                 = "Content-Type"
	ContentTypeJson       ContentType = "application/json"
	ContentTypeJsonUTF8   ContentType = "application/json;charset=UTF-8"
	ContentTypeUrlEncoded ContentType = "application/x-www-form-urlencoded"
)

// HTTPReqBuilder 即可以用于构建 http.Request，也可以用于直接发送请求
// It is not thread-safe, should not be used in multiple goroutines
type HTTPReqBuilder struct {
	ctx                context.Context
	body               io.Reader
	header             http.Header
	query              url.Values
	method, host, path string
	username, password string
}

// NewHTTPReqBuilder returns a new HTTPReqBuilder
func NewHTTPReqBuilder() *HTTPReqBuilder {
	return &HTTPReqBuilder{}
}

func (b *HTTPReqBuilder) reset() {
	b.ctx, b.body = nil, nil
	clear(b.header)
	clear(b.query)
	b.method, b.host, b.path = "", "", ""
	b.username, b.password = "", ""
}

// GetQueryParam 获取 QueryParam
func (b *HTTPReqBuilder) GetQueryParam() url.Values {
	return b.query
}

// SetMethod 设置请求方法
func (b *HTTPReqBuilder) SetMethod(method string) *HTTPReqBuilder {
	b.method = method
	return b
}

// SetBodyString 设置 String 请求体
func (b *HTTPReqBuilder) SetBodyString(body string) *HTTPReqBuilder {
	b.body = bytes.NewReader([]byte(body))
	return b
}

// SetBodyJson 设置 interface 请求体，会自动转为 json
func (b *HTTPReqBuilder) SetBodyJson(body any) *HTTPReqBuilder {
	if body == nil {
		return b
	}
	if bodyBuf, err := sonic.ConfigStd.Marshal(body); err != nil {
		return b
	} else {
		b.body = bytes.NewReader(bodyBuf)
	}
	return b
}

// SetHost 设置请求域名
// 最后是否添加 "/" 均没有影响，会自动去掉
// 比如 SetHost("http://example.com") 和 SetHost("http://example.com/") 效果一样
func (b *HTTPReqBuilder) SetHost(host string) *HTTPReqBuilder {
	b.host = host
	return b
}

// SetPath 设置请求路径，也可以在外层拼接好之后，通过 SetHost 一次设置
// 例如：SetHost("http://example.com").SetPath("/api/v1")
func (b *HTTPReqBuilder) SetPath(path string) *HTTPReqBuilder {
	b.path = path
	return b
}

// SetURL 设置请求 URL，包含 host 和 path
func (b *HTTPReqBuilder) SetURL(url string) *HTTPReqBuilder {
	b.host = url
	return b
}

// SetHeader 设置请求头
func (b *HTTPReqBuilder) SetHeader(key, value string) *HTTPReqBuilder {
	if b.header == nil {
		b.header = make(http.Header)
	}
	b.header.Set(key, value)
	return b
}

// SetHeaderContentType 设置请求头 Content-Type
func (b *HTTPReqBuilder) SetHeaderContentType(contentType ContentType) *HTTPReqBuilder {
	b.SetHeader(ContentTypeHeader, string(contentType))
	return b
}

// AddQueryParams 追加请求参数，如果已经设置过，这个方法会追加
func (b *HTTPReqBuilder) AddQueryParams(key, value string) *HTTPReqBuilder {
	if b.query == nil {
		b.query = make(url.Values)
	}
	b.query.Add(key, value)
	return b
}

// SetQueryParams 设置请求参数，如果已经设置过，这个方法会覆盖
func (b *HTTPReqBuilder) SetQueryParams(key, value string) *HTTPReqBuilder {
	if b.query == nil {
		b.query = make(url.Values)
	}
	b.query.Set(key, value)
	return b
}

func (b *HTTPReqBuilder) SetContext(ctx context.Context) *HTTPReqBuilder {
	b.ctx = ctx
	return b
}

func (b *HTTPReqBuilder) SetUserName(username string) *HTTPReqBuilder {
	b.username = username
	return b
}

func (b *HTTPReqBuilder) SetPassWord(password string) *HTTPReqBuilder {
	b.password = password
	return b
}

func (b *HTTPReqBuilder) build() (*http.Request, error) {
	if b.host == "" {
		return nil, errors.New("host is required")
	}
	if b.method == "" {
		return nil, errors.New("method is required")
	}

	urlTemp := strings.TrimRight(b.host, "/") + b.path
	if b.query != nil && len(b.query) != 0 {
		urlTemp = fmt.Sprintf("%s?%s", urlTemp, b.query.Encode())
	}

	if b.ctx == nil {
		b.ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(b.ctx, b.method, urlTemp, b.body)
	if err != nil {
		return nil, errors.Wrapf(err, "build client NewRequest err")
	}
	req.Header = b.header
	if b.username != "" && b.password != "" {
		req.SetBasicAuth(b.username, b.password)
	}
	log.C(b.ctx).Debugw("HTTPReqBuilder", "request", req)
	return req, nil
}

// Build builds a http.Request，调用后该 HTTPReqBuilder 不应再被使用
func (b *HTTPReqBuilder) Build() (*http.Request, error) {
	return b.build()
}

// Do builds and sends a http.Request，调用后该 HTTPReqBuilder 不应再被使用
func (b *HTTPReqBuilder) Do() (resp *http.Response, err error) {
	req, err := b.build()
	if err != nil {
		return nil, err
	}

	//beginTime := time.Now()
	//defer func() {
	//	metrics.ExtHTTPDurationAndSizeHistogram(b.ctx, beginTime, req, resp)
	//}()
	resp, err = http.DefaultClient.Do(req)
	log.C(b.ctx).Debugw("HTTPReqBuilder", "request", req)
	return resp, err
}

// DoWithTimeout builds and sends a http.Request with timeout，调用后该 HTTPReqBuilder 不应再被使用
func (b *HTTPReqBuilder) DoWithTimeout(timeout time.Duration) (resp *http.Response, err error) {
	req, err := b.build()
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: timeout,
	}

	//beginTime := time.Now()
	//defer func() {
	//	metrics.ExtHTTPDurationAndSizeHistogram(b.ctx, beginTime, req, resp)
	//}()
	resp, err = client.Do(req)
	log.C(b.ctx).Debugw("HTTPReqBuilder DoWithTimeout", "request", req)
	return resp, err
}
