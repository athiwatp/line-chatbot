package shared

import (
	"fmt"
	"net/http"

	"github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/line-chatbot/helper"
)

// Filter model
type Filter struct {
	Page   int
	Limit  int
	Offset int
	Sort   string
	SortBy string
}

// BuildFromHTTPRequest build filter from query params
func (f *Filter) BuildFromHTTPRequest(req *http.Request) *utils.MultiError {
	f.Page = helper.ParseInt(req.URL.Query().Get("page"))
	f.Limit = helper.ParseInt(req.URL.Query().Get("limit"))
	f.SortBy = req.URL.Query().Get("sortBy")
	f.Sort = req.URL.Query().Get("sort")

	multiError := utils.NewMultiError()
	if f.Page < 0 {
		multiError.Append("page", fmt.Errorf("page cannot less than zero"))
	}
	if f.Limit < 0 {
		multiError.Append("limit", fmt.Errorf("limit cannot less than zero"))
	}

	if multiError.HasError() {
		return multiError
	}

	if f.Page == 0 {
		f.Page = 1
	}
	if f.Limit == 0 {
		f.Limit = 10
	}

	return nil
}
