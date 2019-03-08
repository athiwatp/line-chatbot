package usecase

import "github.com/line/line-bot-sdk-go/linebot"

// Usecase abstract interface
type Usecase interface {
	Reply(event *linebot.Event, messages ...string) error
	ProcessMessage(event *linebot.Event, msg *linebot.TextMessage) error
}
