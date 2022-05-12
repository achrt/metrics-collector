package sender

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Request struct {
	Body       interface{}
	Header     http.Header
	RawRequest *http.Request

	timeout int
	url     string
	method  string
	client  *Client
}

const defaultTimeout = 2

func (r *Request) SetHeader(header, value string) *Request {
	r.Header.Set(header, value)
	return r
}

func (r *Request) SetURL(value string) *Request {
	r.url = value
	return r
}

func (r *Request) SetBody(value interface{}) *Request {
	r.Body = value
	return r
}

func (r *Request) SetTimeout(timeout int) *Request {
	r.timeout = timeout
	return r
}

func (r *Request) Post(ctx context.Context, cancel context.CancelFunc) (status int, err error) {
	defer cancel()

	if r.timeout == 0 {
		r.timeout = defaultTimeout
	}

	r.method = http.MethodPost

	ctx2, cancel2 := context.WithTimeout(ctx, time.Duration(r.timeout)*time.Second)
	defer cancel2()

	err = r.rawRequest(ctx2)
	if err != nil {
		return
	}

	resp, err := r.client.execute(r)
	if err != nil {
		return
	}

	<-ctx2.Done()

	status = resp.StatusCode
	return
}

func (r *Request) rawRequest(ctx context.Context) error {
	if r.url == "" || r.method == "" {
		return errors.New("r.URL | r.Method is empty")
	}

	b, err := json.Marshal(r.Body)
	if err != nil {
		return err
	}

	rawReq, err := http.NewRequestWithContext(ctx, r.method, r.url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	rawReq.Header = r.Header
	r.RawRequest = rawReq
	return nil
}
