package usecase

import (
	"github.com/agungdwiprasetyo/line-chatbot/src/entry/domain"
	"github.com/agungdwiprasetyo/line-chatbot/src/entry/repository"
	"github.com/agungdwiprasetyo/line-chatbot/src/shared"
)

type usecaseProfile struct {
	profileRepo repository.Profile
}

func (uc *usecaseProfile) SaveProfile(data *domain.Profile) error {
	return uc.profileRepo.Save(data)
}

func (uc *usecaseProfile) FindAllProfile(filter *shared.Filter) (result shared.Result) {
	filter.Offset = (filter.Page - 1) * filter.Limit

	repoRes := uc.profileRepo.FindAll(filter)
	if repoRes.Error != nil {
		result.Error = repoRes.Error
		return
	}

	result.Data = repoRes.Data
	result.Total = uc.profileRepo.Count(filter)
	return
}

func (uc *usecaseProfile) FindProfileByID(id string) (res shared.Result) {
	return
}

func (uc *usecaseProfile) Update(data *domain.Profile) (res shared.Result) {
	return
}

func (uc *usecaseProfile) Remove(id string) (res shared.Result) {
	return
}
