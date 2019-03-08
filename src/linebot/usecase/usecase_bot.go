package usecase

import (
	"fmt"

	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/repository"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
	"github.com/line/line-bot-sdk-go/linebot"
)

type usecaseImpl struct {
	bot             *linebot.Client
	repository      *shared.Repository
	eventRepository repository.Event
}

// NewUsecase constructor
func NewUsecase(repo *shared.Repository, bot *linebot.Client) Usecase {
	uc := new(usecaseImpl)
	uc.bot = bot
	uc.repository = repo
	uc.eventRepository = repository.NewRepositoryEventMongo(repo)

	return uc
}

func (uc *usecaseImpl) Reply(event *linebot.Event, messages ...string) error {
	var lineMessages []linebot.SendingMessage
	for _, msg := range messages {
		lineMessages = append(lineMessages, linebot.NewTextMessage(msg))
	}

	if _, err := uc.bot.ReplyMessage(event.ReplyToken, lineMessages...).Do(); err != nil {
		return err
	}

	return nil
}

func (uc *usecaseImpl) ProcessMessage(event *linebot.Event, msg *linebot.TextMessage) error {
	go func() {
		var e domain.Event
		e.Build(event)
		uc.eventRepository.Save(&e)
	}()

	return uc.Reply(event, "(under construction)", fmt.Sprintf("Mantul gan => %s", msg.Text))
}
