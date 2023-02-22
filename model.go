package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	XRequestIDKey = "X-Request-ID"
)

type Request struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body,omitempty"`
}

func (r *Request) parseBody(src io.ReadCloser) error {
	defer src.Close()
	body, err := ioutil.ReadAll(src)
	if err != nil {
		return fmt.Errorf("reading body: %w", err)
	}
	return json.Unmarshal(body, r)
}

func (r *Request) validate() error {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		return fmt.Errorf("method not allowed")
	}
	_, err := url.Parse(r.Url)
	if err != nil {
		return fmt.Errorf("url not valid")
	}
	return nil
}

func (r *Request) marshaling() string {
	body, _ := json.Marshal(r)
	return string(body)
}

type Response struct {
	ID      string              `json:"id"`
	Status  string              `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
	Body    json.RawMessage     `json:"body,omitempty"`
}

func (r *Response) marshaling() string {
	body, _ := json.Marshal(r)
	return string(body)
}
