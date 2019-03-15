package middleware

import (
	"net/http"

	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
)

// basicAuth middleware
type basicAuth struct {
	username, password string
}

// newBasicAuth constructor
func newBasicAuth(username, password string) *basicAuth {
	return &basicAuth{username, password}
}

// BasicAuth implement middlewarefunc from mux
func (b *basicAuth) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != b.username || pass != b.password {
			response := shared.NewHTTPResponse(http.StatusUnauthorized, "invalid signature")
			response.JSON(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
