package main

import "sync"

type Storage interface {
	Add(req *Request, resp *Response) error
}

func NewStorage() Storage {
	return &storage{
		requests:  make(map[string]string),
		responses: make(map[string]string),
	}
}

type storage struct {
	lock      sync.Mutex
	requests  map[string]string
	responses map[string]string
}

func (s *storage) Add(req *Request, resp *Response) error {
	defer s.lock.Unlock()
	reqStr := req.marshaling()
	respStr := resp.marshaling()
	s.lock.Lock()
	s.requests[resp.ID] = reqStr
	s.responses[resp.ID] = respStr
	return nil
}
