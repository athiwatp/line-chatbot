package domain

import (
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/mgo.v2/bson"
)

// Event domain model
type Event struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	ReplyToken string        `bson:"reply_token" json:"replyToken"`
	Type       string        `bson:"type" json:"type"`
	Timestamp  time.Time     `bson:"timestamp" json:"timestamp"`
	Source     struct {
		GroupID string `bson:"group_id" json:"groupID"`
		Type    string `bson:"type" json:"type"`
		UserID  string `bson:"user_id" json:"userID"`
	} `bson:"source" json:"source"`
	Message struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Text string `json:"text"`
	} `bson:"message" json:"message"`
}

// Build event domain from line event bot data
func (e *Event) Build(lineBotData *linebot.Event) {
	e.ReplyToken = lineBotData.ReplyToken
	e.Type = string(lineBotData.Type)
	e.Timestamp = lineBotData.Timestamp

	if lineBotData.Source != nil {
		e.Source.GroupID = lineBotData.Source.GroupID
		e.Source.Type = string(lineBotData.Source.Type)
		e.Source.UserID = lineBotData.Source.UserID
	}
}
