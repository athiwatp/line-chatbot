package delivery

import (
	"net/http"

	"github.com/agungdwiprasetyo/go-line-chatbot/middleware"
	entryUseCase "github.com/agungdwiprasetyo/go-line-chatbot/src/entry/usecase"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
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

// Mount router to entry handler
func (h *Handler) Mount(root string) {
	http.Handle("/"+root+"/events", middleware.BasicAuth(http.HandlerFunc(h.findAllEventLog)))
	http.Handle("/"+root+"/clearlog", middleware.BasicAuth(http.HandlerFunc(h.clearAllLog)))
	http.Handle("/"+root+"/users", middleware.BasicAuth(http.HandlerFunc(h.findAllUser)))
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
