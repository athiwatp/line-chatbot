package repository

import (
	"github.com/agungdwiprasetyo/line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
)

// Entry abstract interface
type Entry interface {
	Create(data *domain.Entry) error
}

// Event interface
type Event interface {
	Count(filter *shared.Filter) int
	FindAll(filter *shared.Filter) ([]*domain.Event, error)
	Save(data *domain.Event) error
	RemoveAll() error
}

// Profile interface
type Profile interface {
	Count(filter *shared.Filter) int
	Save(*domain.Profile) error
	FindAll(filter *shared.Filter) (res shared.Result)
	FindByID(string) shared.Result
}
