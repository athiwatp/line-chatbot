package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agungdwiprasetyo/line-chatbot/middleware"
	"github.com/gorilla/mux"
)

func (s *Service) ServeHTTP() {
	router := mux.NewRouter()

	authorization := middleware.NewAuthorization(s.conf)

	// mount main linebot handler
	s.linebotHandler.Mount(router, authorization)

	// mount entry rest handler
	entryRouter := router.PathPrefix("/entry").Subrouter()
	entryRouter.Use(authorization.BasicAuth)
	s.entryHandler.Mount(entryRouter)

	port := fmt.Sprintf(":%s", s.conf.HTTPPort)
	log.Println("Server running on port", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
