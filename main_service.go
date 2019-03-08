package main

import (
	"log"

	"github.com/agungdwiprasetyo/go-line-chatbot/config"
	linebotHandler "github.com/agungdwiprasetyo/go-line-chatbot/src/linebot/delivery"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Service main
type Service struct {
	conf           *config.Config
	linebotHandler *linebotHandler.Handler
}

func initMainService(conf *config.Config) *Service {
	bot, err := linebot.New(conf.LineChannelSecret, conf.LineChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	linebotHandler := linebotHandler.NewHandler(bot)

	return &Service{
		conf:           conf,
		linebotHandler: linebotHandler,
	}
}
