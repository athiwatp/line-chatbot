package usecase

import (
	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/entry/repository"
	"github.com/agungdwiprasetyo/go-line-chatbot/src/shared"
)

type usecaseEvent struct {
	repository      *shared.Repository
	eventRepository repository.Event
}

func (uc *usecaseEvent) SaveLogEvent(data *domain.Event) error {
	return uc.eventRepository.Save(data)
}

func (uc *usecaseEvent) FindAllEvent(filter *shared.Filter) ([]*domain.Event, error) {
	filter.Offset = (filter.Page - 1) * filter.Limit
	return uc.eventRepository.FindAll(filter)
}

func (uc *usecaseEvent) ClearAllEventLog() error {
	return uc.eventRepository.RemoveAll()
}
