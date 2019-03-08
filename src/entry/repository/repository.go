package repository

import "github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"

// Entry abstract interface
type Entry interface {
	Create(data *domain.Entry) error
}

// Event interface
type Event interface {
	FindAll() ([]*domain.Event, error)
	Save(data *domain.Event) error
}
