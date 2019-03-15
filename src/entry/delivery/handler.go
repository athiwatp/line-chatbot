package delivery

import (
	"net/http"

	entryUseCase "github.com/agungdwiprasetyo/line-chatbot/src/entry/usecase"
	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
	"github.com/gorilla/mux"
)

// Handler model
type Handler struct {
	entryUsecase entryUseCase.Usecase
}

// NewHandler constructor
func NewHandler(entryUsecase entryUseCase.Usecase) *Handler {
	return &Handler{
		entryUsecase: entryUsecase,
	}
}

// Mount router to entry handler (prefix => "/entry") with basic authorization
func (h *Handler) Mount(entry *mux.Router) {
	entry.HandleFunc("/events", h.findAllEventLog)
	entry.HandleFunc("/clearlog", h.clearAllLog)
	entry.HandleFunc("/users", h.findAllUser)
}

func (h *Handler) findAllEventLog(w http.ResponseWriter, req *http.Request) {
	var filter shared.Filter
	if err := filter.BuildFromHTTPRequest(req); err != nil {
		response := shared.NewHTTPResponse(http.StatusBadRequest, "failed to validate params", err)
		response.JSON(w)
		return
	}

	result := h.entryUsecase.FindAllEvent(&filter)
	if result.Error != nil {
		response := shared.NewHTTPResponse(http.StatusInternalServerError, result.Error.Error())
		response.JSON(w)
		return
	}

	meta := shared.Meta{Page: filter.Page, Limit: filter.Limit, TotalRecords: result.Total}
	response := shared.NewHTTPResponse(http.StatusOK, "success", result.Data, meta)
	response.JSON(w)
}

func (h *Handler) clearAllLog(w http.ResponseWriter, req *http.Request) {
	err := h.entryUsecase.ClearAllEventLog()
	if err != nil {
		response := shared.NewHTTPResponse(http.StatusInternalServerError, err.Error())
		response.JSON(w)
		return
	}

	response := shared.NewHTTPResponse(http.StatusOK, "success")
	response.JSON(w)
}

func (h *Handler) findAllUser(w http.ResponseWriter, req *http.Request) {
	var filter shared.Filter
	if err := filter.BuildFromHTTPRequest(req); err != nil {
		response := shared.NewHTTPResponse(http.StatusBadRequest, "failed to validate params", err)
		response.JSON(w)
		return
	}

	result := h.entryUsecase.FindAllProfile(&filter)
	if result.Error != nil {
		response := shared.NewHTTPResponse(http.StatusInternalServerError, result.Error.Error())
		response.JSON(w)
		return
	}

	meta := shared.Meta{Page: filter.Page, Limit: filter.Limit, TotalRecords: result.Total}
	response := shared.NewHTTPResponse(http.StatusOK, "success", result.Data, meta)
	response.JSON(w)
}
