package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type Service interface {
	Send(id string, req *Request) (*Response, error)
}

type service struct {
	cli Client
	db  Storage
}

func NewService() Service {
	return &service{
		cli: NewClient(),
		db:  NewStorage(),
	}
}

func (s *service) Send(id string, req *Request) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cliResp, err := s.cli.SendRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("client: send request: %w", err)
	}

	resp := Response{
		ID:      id,
		Status:  cliResp.Status,
		Headers: cliResp.Header,
	}
	l := cliResp.Header.Get("Content-Length")
	if l != "" {
		n, err := strconv.Atoi(l)
		if err == nil {
			resp.Length = n
		}
	}
	// resp.Body, _ = ioutil.ReadAll(cliResp.Body)
	if err = s.db.Add(req, &resp); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	return &resp, nil
}
