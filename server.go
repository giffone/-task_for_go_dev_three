package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	cfg *conf
	srv *http.Server
}

func NewServer(cfg *conf) *server {
	return &server{
		srv: &http.Server{
			Addr:         cfg.addr,
			WriteTimeout: cfg.write,
			ReadTimeout:  cfg.read,
			Handler:      registerHandlers(),
		},
		cfg: cfg,
	}
}

func (s *server) Run() error {
	log.Printf("https://localhost%s is listening...\n", s.cfg.addr)
	if err := s.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("listenAndServe: %w", err)
	}
	return nil
}
