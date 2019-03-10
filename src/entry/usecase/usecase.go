package usecase

import (
	"github.com/agungdwiprasetyo/line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/line-chatbot/src/entry/repository"
	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
)

// Usecase abstract interface
type Usecase interface {
	SaveLogEvent(data *domain.Event) error
	FindAllEvent(filter *shared.Filter) shared.Result
	ClearAllEventLog() error
	SaveProfile(data *domain.Profile) error
	FindAllProfile(filter *shared.Filter) shared.Result
	FindProfileByID(id string) shared.Result
}

type usecaseImpl struct {
	*usecaseEvent
	*usecaseProfile
}

// NewUsecase constructor
func NewUsecase(repo *shared.Repository) Usecase {
	eventUsecase := new(usecaseEvent)
	eventUsecase.repository = repo
	eventUsecase.eventRepository = repository.NewRepositoryEventMongo(repo)

	profileUsecase := new(usecaseProfile)
	profileUsecase.profileRepo = repository.NewRepositoryProfileMongo(repo)

	return &usecaseImpl{usecaseEvent: eventUsecase, usecaseProfile: profileUsecase}
}
