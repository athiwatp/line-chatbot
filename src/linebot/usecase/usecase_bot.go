package usecase

import (
	"fmt"

	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/repository"
	lineBotDomain "github.com/agungdwiprasetyo/go-line-chatbot/src/linebot/domain"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
	lineService "github.com/agungdwiprasetyo/go-line-chatbot/src/shared/service/line"
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

func (uc *usecaseImpl) PushMessageToChannel(to, title, message string) error {
	var lineMessage lineBotDomain.Message

	lineMessage.To = to
	lineMessage.Messages = append(lineMessage.Messages, lineBotDomain.ContentMessage{
		Type: "flex", AltText: title, Contents: lineBotDomain.ContentFormat{
			Type: "bubble", Body: lineBotDomain.ContentBody{
				Type: "box", Layout: "horizontal", Contents: []lineBotDomain.Content{
					lineBotDomain.Content{
						Type: "text", Text: message,
					},
				},
			},
		},
	})

	return lineService.SendMessage(&lineMessage)
}
