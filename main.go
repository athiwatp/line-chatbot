package main

import (
	"log"
	"sync"

	"github.com/agungdwiprasetyo/line-chatbot/config"
	env "github.com/joho/godotenv"
)

func main() {
	if err := env.Load(".env"); err != nil {
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
