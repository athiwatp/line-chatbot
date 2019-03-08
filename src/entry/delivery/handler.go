package delivery

import (
	"net/http"

	entryUseCase "github.com/agungdwiprasetyo/go-line-chatbot/src/entry/usecase"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
)

type Handler struct {
	entryUsecase entryUseCase.Usecase
}

func NewHandler(entryUsecase entryUseCase.Usecase) *Handler {
	return &Handler{
		entryUsecase: entryUsecase,
	}
}

func (h *Handler) Mount() {
	http.HandleFunc("/entry/events", h.findAllEventLog)
	http.HandleFunc("/entry/clearlog", h.clearAllLog)
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

func (h *Handler) clearAllLog(w http.ResponseWriter, req *http.Request) {
	err := h.entryUsecase.ClearAllEventLog()
	if err != nil {
		response := shared.NewHTTPResponse(http.StatusInternalServerError, err.Error())
		response.JSON(w)
	}

	response := shared.NewHTTPResponse(http.StatusOK, "success")
	response.JSON(w)
}
