package middleware

import (
	"net/http"
	"os"

	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
)

// BasicAuth middleware
func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey != os.Getenv("API_KEY") {
			response := shared.NewHTTPResponse(http.StatusUnauthorized, "invalid signature")
			response.JSON(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
