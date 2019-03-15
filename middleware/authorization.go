package middleware

import "github.com/agungdwiprasetyo/line-chatbot/config"

// Authorization base model for all type of authorization
type Authorization struct {
	*basicAuth
	*bearer
	*apiKey
}

// NewAuthorization construct new middleware for authorization
func NewAuthorization(conf *config.Config) *Authorization {
	return &Authorization{
		basicAuth: newBasicAuth(conf.Key.Username, conf.Key.Password),
		bearer:    newBearer(conf.Key.PublicKey, conf.Key.PrivateKey),
		apiKey:    newAPIKey(conf.Key.APIKey),
	}
}
