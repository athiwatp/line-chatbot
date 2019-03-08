package delivery

import (
	"log"
	"net/http"

	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Handler struct {
	bot *linebot.Client
}

func NewHandler(lineClient *linebot.Client) *Handler {
	return &Handler{bot: lineClient}
}

func (h *Handler) Mount() {
	http.HandleFunc("/callback", h.Callback)
}

func (h *Handler) Callback(w http.ResponseWriter, req *http.Request) {
	events, err := h.bot.ParseRequest(req)
	if err != nil {
		var code int
		if err == linebot.ErrInvalidSignature {
			code = http.StatusBadRequest
		} else {
			code = http.StatusInternalServerError
		}
		response := shared.NewHTTPResponse(code, err.Error())
		response.JSON(w)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = h.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}

	response := shared.NewHTTPResponse(http.StatusOK, "ok")
	response.JSON(w)
}
