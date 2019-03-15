package middleware

import (
	"net/http"

	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
)

// apiKey middleware with x-api-key
type apiKey struct {
	apiKey string
}

// newAPIKey constructor
func newAPIKey(key string) *apiKey {
	return &apiKey{key}
}

// APIKey middleware with api key
func (a *apiKey) APIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey != a.apiKey {
			response := shared.NewHTTPResponse(http.StatusUnauthorized, "invalid signature")
			response.JSON(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
