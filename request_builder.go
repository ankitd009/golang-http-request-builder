package httpreqbuilder

import (
	"context"
	"io"
	"net/http"
	"time"
)

var (
	DefaultTimeoutInSeconds = 10
)

type ReqBuilder struct {
	methodType  string
	url         string
	headers     map[string]string
	body        io.Reader
	queryParams map[string]string
	cookies     []*http.Cookie
	timeout     int
}

func SetDefaultTimeout(timeout int) {
	DefaultTimeoutInSeconds = timeout
}

func New(methodType, url string) *ReqBuilder {
	return &ReqBuilder{
		methodType:  methodType,
		url:         url,
		headers:     make(map[string]string),
		queryParams: make(map[string]string),
		cookies:     make([]*http.Cookie, 0, 0),
		body:        nil,
	}
}

func (rb *ReqBuilder) WithCookie(cookie *http.Cookie) *ReqBuilder {
	cookies := rb.cookies
	cookies = append(cookies, cookie)
	rb.cookies = cookies
	return rb
}

func (rb *ReqBuilder) WithHeader(name, value string) *ReqBuilder {
	rb.headers[name] = value
	return rb
}

func (rb *ReqBuilder) WithBody(reader io.Reader) *ReqBuilder {
	rb.body = reader
	return rb
}

func (rb *ReqBuilder) WithQueryParam(name, value string) *ReqBuilder {
	rb.queryParams[name] = value
	return rb
}

func (rb *ReqBuilder) WithTimeout(timeout int) *ReqBuilder {
	rb.timeout = timeout
	return rb
}

func (rb *ReqBuilder) Build() (*http.Request, context.CancelFunc, error) {

	timeout := rb.timeout
	if rb.timeout <= 0 {
		timeout = DefaultTimeoutInSeconds
	}

	req, err := http.NewRequest(rb.methodType, rb.url, rb.body)
	if err != nil {
		return nil, nil, err
	}

	// return cancel function from here
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	req = req.WithContext(ctx)

	for k, v := range rb.headers {
		req.Header.Set(k, v)
	}

	for _, cookie := range rb.cookies {
		req.AddCookie(cookie)
	}

	q := req.URL.Query()
	for k, v := range rb.queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return req, cancel, nil
}
