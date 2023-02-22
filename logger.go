package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type (
	handlerF func(w http.ResponseWriter, r *http.Request) error
	jsonF    func(w http.ResponseWriter, status int, respBody any) error
)

type Middleware interface {
	Logger(handler handlerF) http.Handler
	JSON(w http.ResponseWriter, status int, respBody any) error
}

func NewMiddleware() Middleware {
	return &middleware{}
}

type middleware struct{}

func (m *middleware) Logger(handler handlerF) http.Handler {
	return &logger{
		handler: handler,
		json:    m.JSON,
	}
}

type logger struct {
	handler handlerF
	json    jsonF
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewString()
	var err error

	defer func(t time.Time) {
		log.Printf("{id: %s}{method: %s}{path: %s}{err: %s}{bench msec: %v}\n",
			id, r.Method, r.URL.Path, err, time.Since(t).Milliseconds())
	}(time.Now())

	r.Header.Set(XRequestIDKey, id)

	if r.Method != http.MethodPost {
		l.json(w, http.StatusMethodNotAllowed, Response{ID: id})
		return
	}

	if err = l.handler(w, r); err != nil {
		if errors.Is(err, ErrService) {
			l.json(w, http.StatusInternalServerError, Response{ID: id})
			return
		}
		if errors.Is(err, context.DeadlineExceeded) {
			l.json(w, http.StatusGatewayTimeout, Response{ID: id})
			return
		}
		l.json(w, http.StatusBadRequest, Response{ID: id})
	}
}

func (m *middleware) JSON(w http.ResponseWriter, status int, respBody any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(respBody)
}
