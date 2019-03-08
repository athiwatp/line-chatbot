package repository

import "github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"

// Entry abstract interface
type Entry interface {
	Create(data *domain.Entry) error
}
