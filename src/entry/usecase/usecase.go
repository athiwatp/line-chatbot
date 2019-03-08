package usecase

import (
	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/repository"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
)

// Usecase abstract interface
type Usecase interface {
	SaveLogEvent(data *domain.Event) error
	FindAllEvent() ([]*domain.Event, error)
}

type usecaseImpl struct {
	*usecaseEvent
}

// NewUsecase constructor
func NewUsecase(repo *shared.Repository) Usecase {
	eventUsecase := new(usecaseEvent)
	eventUsecase.repository = repo
	eventUsecase.eventRepository = repository.NewRepositoryEventMongo(repo)

	return &usecaseImpl{eventUsecase}
}
