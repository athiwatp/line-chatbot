package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"
	entryUseCase "github.com/agungdwiprasetyo/go-line-chatbot/src/entry/usecase"
	botUseCase "github.com/agungdwiprasetyo/go-line-chatbot/src/linebot/usecase"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
	"github.com/agungdwiprasetyo/go-utils/debug"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Handler struct {
	bot          *linebot.Client
	botUsecase   botUseCase.Usecase
	entryUsecase entryUseCase.Usecase
}

func NewHandler(lineClient *linebot.Client, botUsecase botUseCase.Usecase, entryUsecase entryUseCase.Usecase) *Handler {
	return &Handler{
		bot:          lineClient,
		botUsecase:   botUsecase,
		entryUsecase: entryUsecase,
	}
}

func (h *Handler) Mount() {
	http.HandleFunc("/callback", h.callback)
	http.HandleFunc("/events", h.findAllEventLog)
}

func (h *Handler) callback(w http.ResponseWriter, req *http.Request) {
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
		debug.PrintJSON(event)
		if event.Type == linebot.EventTypeMessage {
			var e domain.Event
			jsonStr, _ := json.Marshal(event)
			e.Build(jsonStr)
			go h.entryUsecase.SaveLogEvent(&e)

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyMessage := h.botUsecase.ProcessMessage(message.Text)
				if _, err = h.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					log.Print(err)
				}
			}

			profile := h.bot.GetProfile(event.Source.UserID)
			debug.PrintJSON(profile)
		}
	}

	response := shared.NewHTTPResponse(http.StatusOK, "ok")
	response.JSON(w)
}

func (h *Handler) findAllEventLog(w http.ResponseWriter, req *http.Request) {
	data, err := h.entryUsecase.FindAllEvent()
	if err != nil {
		response := shared.NewHTTPResponse(http.StatusInternalServerError, err.Error())
		response.JSON(w)
	}

	response := shared.NewHTTPResponse(http.StatusOK, "success", data)
	response.JSON(w)
}
