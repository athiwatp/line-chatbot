package shared

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"reflect"

	utils "github.com/agungdwiprasetyo/go-utils"
)

// HTTPResponse abstract interface
type HTTPResponse interface {
	JSON(w http.ResponseWriter)
	XML(w http.ResponseWriter)
}

type (
	// httpResponse model
	httpResponse struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
	}
)

// NewHTTPResponse for create common response, data must in first params and meta in second params
func NewHTTPResponse(code int, message string, params ...interface{}) HTTPResponse {
	commonResponse := new(httpResponse)

	for _, param := range params {
		refValue := reflect.ValueOf(param)
		if refValue.Kind() == reflect.Ptr {
			refValue = refValue.Elem()
		}
		param = refValue.Interface()

		switch param.(type) {
		case utils.MultiError:
			multiError := param.(utils.MultiError)
			commonResponse.Errors = multiError.ToMap()
		default:
			commonResponse.Data = param
		}
	}

	if code < 400 {
		commonResponse.Success = true
	}

	commonResponse.Code = code
	commonResponse.Message = message
	return commonResponse
}

// JSON for set http JSON response (Content-Type: application/json)
func (resp *httpResponse) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// XML for set http XML response (Content-Type: application/xml)
func (resp *httpResponse) XML(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(resp.Code)
	xml.NewEncoder(w).Encode(resp)
}
