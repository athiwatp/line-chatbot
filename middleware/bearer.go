package middleware

import (
	"crypto/rsa"
	"net/http"
)

type bearer struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func newBearer(publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) *bearer {
	return &bearer{publicKey, privateKey}
}

// Bearer implement middlewarefunc from mux
func (b *bearer) Bearer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: add jwt verify
		next.ServeHTTP(w, r)
	})
}
