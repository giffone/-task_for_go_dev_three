package main

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrService = errors.New("service")

type handlers struct {
	svc Service
	mid Middleware
}

func registerHandlers() http.Handler {
	h := handlers{
		svc: NewService(),
		mid: NewMiddleware(),
	}
	mux := http.NewServeMux()
	mux.Handle("/proxy", h.mid.Logger(h.proxy))
	return mux
}

func (h *handlers) proxy(w http.ResponseWriter, r *http.Request) error {
	req := Request{}
	id := r.Header.Get(XRequestIDKey)
	if err := req.parseBody(r.Body); err != nil {
		return err
	}
	if err := req.validate(); err != nil {
		return err
	}
	resp, err := h.svc.Send(id, &req)
	if err != nil {
		return fmt.Errorf("%w: send: %w", ErrService, err)
	}
	return h.mid.JSON(w, http.StatusOK, resp)
}
