package usecase

import (
	"fmt"

	"github.com/agungdwiprasetyo/line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/line-chatbot/src/entry/repository"
	lineBotDomain "github.com/agungdwiprasetyo/line-chatbot/src/linebot/domain"
	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
	lineService "github.com/agungdwiprasetyo/line-chatbot/src/shared/service/line"
	"github.com/line/line-bot-sdk-go/linebot"
)

type usecaseImpl struct {
	bot               *linebot.Client
	repository        *shared.Repository
	eventRepository   repository.Event
	profileRepository repository.Profile
}

// NewUsecase constructor
func NewUsecase(repo *shared.Repository, bot *linebot.Client) Usecase {
	uc := new(usecaseImpl)
	uc.bot = bot
	uc.repository = repo
	uc.eventRepository = repository.NewRepositoryEventMongo(repo)
	uc.profileRepository = repository.NewRepositoryProfileMongo(repo)

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
	// TODO: add AI to process output from incoming message
	return uc.Reply(event, "(under construction)", fmt.Sprintf("Mantul gan, masih brute force :) => %s", msg.Text))
}

func (uc *usecaseImpl) SaveLog(event *linebot.Event) error {
	var e domain.Event
	e.ReplyToken = event.ReplyToken
	e.Type = string(event.Type)
	e.Timestamp = event.Timestamp

	message, ok := event.Message.(*linebot.TextMessage)
	if ok && message != nil {
		e.Message.ID = message.ID
		e.Message.Text = message.Text
	}

	if event.Source != nil {
		profile, err := uc.bot.GetProfile(event.Source.UserID).Do()
		if err != nil {
			profile = &linebot.UserProfileResponse{DisplayName: "Agung DP"}
		}

		var source domain.Profile
		source.ID = event.Source.UserID
		source.Type = string(event.Source.Type)
		source.Name = profile.DisplayName
		source.Avatar = profile.PictureURL
		source.StatusMessage = profile.StatusMessage

		e.SourceID = source.ID
		e.SourceType = source.Type

		if err := uc.profileRepository.Save(&source); err != nil {
			return err
		}
	}

	if err := uc.eventRepository.Save(&e); err != nil {
		return err
	}

	return nil
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
