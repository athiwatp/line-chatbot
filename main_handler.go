package main

import (
	"log"
	"net/http"
)

func (s *Service) ServeHTTP() {
	s.linebotHandler.Mount()
	s.entryHandler.Mount("entry")

	if err := http.ListenAndServe(":"+s.conf.HTTPPort, nil); err != nil {
		log.Fatal(err)
	}
}
