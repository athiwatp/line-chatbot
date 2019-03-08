package delivery

import (
	"fmt"
	"net/http"

	"github.com/agungdwiprasetyo/go-line-chatbot/helper"
	"github.com/agungdwiprasetyo/go-line-chatbot/middleware"
	entryUseCase "github.com/agungdwiprasetyo/go-line-chatbot/src/entry/usecase"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
	"github.com/agungdwiprasetyo/go-utils"
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
	http.Handle("/entry/events", middleware.BasicAuth(http.HandlerFunc(h.findAllEventLog)))
	http.Handle("/entry/clearlog", middleware.BasicAuth(http.HandlerFunc(h.clearAllLog)))
}

func (h *Handler) findAllEventLog(w http.ResponseWriter, req *http.Request) {
	filter := shared.Filter{
		Page:   helper.ParseInt(req.URL.Query().Get("page")),
		Limit:  helper.ParseInt(req.URL.Query().Get("limit")),
		SortBy: req.URL.Query().Get("sortBy"),
		Sort:   req.URL.Query().Get("sort"),
	}

	multiError := utils.NewMultiError()
	if filter.Page < 0 {
		multiError.Append("page", fmt.Errorf("page cannot less than zero"))
	}
	if filter.Limit < 0 {
		multiError.Append("limit", fmt.Errorf("limit cannot less than zero"))
	}

	if !multiError.IsNil() {
		response := shared.NewHTTPResponse(http.StatusBadRequest, "failed to validate params", multiError)
		response.JSON(w)
		return
	}

	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Limit == 0 {
		filter.Limit = 10
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
