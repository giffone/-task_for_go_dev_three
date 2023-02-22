package main

import (
	"flag"
	"strconv"
	"time"
)

type conf struct {
	addr  string
	read  time.Duration
	write time.Duration
}

func NewConfig() (*conf, error) {
	var c conf
	flag.StringVar(&c.addr, "addr", ":8000", "port address in format \":8080\"")
	var read, write string
	flag.StringVar(&read, "read", "5", "server read duration")
	flag.StringVar(&write, "write", "10", "server write duration")

	r, err := strconv.ParseInt(read, 10, 64)
	if err != nil {
		return nil, err
	}
	c.read = time.Duration(r) * time.Second

	w, err := strconv.ParseInt(write, 10, 64)
	if err != nil {
		return nil, err
	}
	c.write = time.Duration(w) * time.Second
	return &c, nil
}
