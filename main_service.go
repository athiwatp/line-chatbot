package main

import (
	"log"

	"github.com/agungdwiprasetyo/go-line-chatbot/config"
	entryUseCase "github.com/agungdwiprasetyo/go-line-chatbot/src/entry/usecase"
	linebotHandler "github.com/agungdwiprasetyo/go-line-chatbot/src/linebot/delivery"
	botUseCase "github.com/agungdwiprasetyo/go-line-chatbot/src/linebot/usecase"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Service main
type Service struct {
	conf           *config.Config
	linebotHandler *linebotHandler.Handler
}

func initMainService(conf *config.Config) *Service {
	repository := shared.NewRepository(conf.MongoDB)

	bot, err := linebot.New(conf.LineChannelSecret, conf.LineChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	botUsecase := botUseCase.NewUsecase(repository)
	entryUsecase := entryUseCase.NewUsecase(repository)

	linebotHandler := linebotHandler.NewHandler(bot, botUsecase, entryUsecase)

	return &Service{
		conf:           conf,
		linebotHandler: linebotHandler,
	}
}
