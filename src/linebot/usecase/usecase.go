package usecase

import (
	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/repository"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
)

// Usecase abstract interface
type Usecase interface {
	ProcessMessage(message string) string
}

type usecaseImpl struct {
	repository      *shared.Repository
	eventRepository repository.Event
}

// NewUsecase constructor
func NewUsecase(repo *shared.Repository) Usecase {
	uc := new(usecaseImpl)
	uc.repository = repo
	uc.eventRepository = repository.NewRepositoryEventMongo(repo)

	return uc
}

func (uc *usecaseImpl) ProcessMessage(message string) string {
	return "Buatan agung => " + message
}
