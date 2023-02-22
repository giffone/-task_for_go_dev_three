package main

import (
	"log"
)

func main() {
	conf, err := NewConfig()
	if err != nil {
		log.Fatalf("config: %s", err.Error())
	}

	s := NewServer(conf)
	if err := s.Run(); err != nil {
		log.Fatalf("run server: %s", err.Error())
	}
}
