package main

import (
	"log"
	"sync"

	"github.com/agungdwiprasetyo/go-line-chatbot/config"
	env "github.com/joho/godotenv"
)

func main() {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}

	conf := config.Init()
	service := initMainService(conf)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.ServeHTTP()
	}()

	wg.Wait()
}
