package sender

import (
	"errors"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) SetTransport(maxIdleConns int) *Client {
	c.httpClient.Transport = &http.Transport{MaxIdleConns: maxIdleConns}
	return c
}

func (c *Client) R() *Request {
	return &Request{
		Header: http.Header{},
		client: c,
	}
}

func (c *Client) execute(req *Request) (*Response, error) {
	if req.RawRequest == nil {
		return nil, errors.New("req.RawRequest is nil")
	}
	resp, err := c.httpClient.Do(req.RawRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := &Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Proto:      resp.Proto,
		Body:       resp.Body,
	}

	return r, err
}
