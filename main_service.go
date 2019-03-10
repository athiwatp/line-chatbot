package main

import (
	"log"

	"github.com/agungdwiprasetyo/line-chatbot/config"
	entryhandler "github.com/agungdwiprasetyo/line-chatbot/src/entry/delivery"
	entryUseCase "github.com/agungdwiprasetyo/line-chatbot/src/entry/usecase"
	linebothandler "github.com/agungdwiprasetyo/line-chatbot/src/linebot/delivery"
	botUseCase "github.com/agungdwiprasetyo/line-chatbot/src/linebot/usecase"
	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Service main
type Service struct {
	conf           *config.Config
	linebotHandler *linebothandler.Handler
	entryHandler   *entryhandler.Handler
}

func initMainService(conf *config.Config) *Service {
	repository := shared.NewRepository(conf.MongoDB)

	bot, err := linebot.New(conf.LineChannelSecret, conf.LineChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	botUsecase := botUseCase.NewUsecase(repository, bot)
	entryUsecase := entryUseCase.NewUsecase(repository)

	linebotHandler := linebothandler.NewHandler(bot, botUsecase, entryUsecase)
	entryHandler := entryhandler.NewHandler(entryUsecase)

	return &Service{
		conf:           conf,
		linebotHandler: linebotHandler,
		entryHandler:   entryHandler,
	}
}
