package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type Client interface {
	SendRequest(ctx context.Context, req *Request) (*http.Response, error)
}

func NewClient() Client {
	return &client{
		cli: http.Client{
			Timeout: 0,
		},
	}
}

type client struct {
	cli http.Client
}

func (c *client) SendRequest(ctx context.Context, req *Request) (*http.Response, error) {

	b := bytes.NewReader(req.Body)

	r, err := http.NewRequestWithContext(ctx, req.Method, req.Url, b)
	if err != nil {
		return nil, fmt.Errorf("make new request: %w", err)
	}
	for k, v := range req.Headers {
		r.Header.Set(k, v)
	}
	resp, err := c.cli.Do(r)
	if err != nil {
		return nil, fmt.Errorf("request do: %w", err)
	}
	return resp, nil
}
