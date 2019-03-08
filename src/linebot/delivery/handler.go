package delivery

import (
	"fmt"
	"net/http"

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

		profile, err := h.bot.GetProfile(event.Source.UserID).Do()
		if err != nil {
			profile = &linebot.UserProfileResponse{DisplayName: "Agung DP"}
		}
		debug.PrintJSON(profile)

		switch event.Type {
		case linebot.EventTypeJoin:
			h.botUsecase.Reply(event, fmt.Sprintf("Hello %s :)", profile.DisplayName))

		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				h.botUsecase.ProcessMessage(event, message)
			}

		}
	}

	response := shared.NewHTTPResponse(http.StatusOK, "ok")
	response.JSON(w)
}
