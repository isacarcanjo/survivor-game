package usecases

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type UpdateUserCase struct {
	repo repositories.ISurvivorRepository
}

func NewService(r repositories.ISurvivorRepository) *UpdateUserCase {
	return &UpdateUserCase{
		repo: r,
	}
}

func (s *UpdateUserCase) Execute(id string, input UpdateSurvivorInput) (entitys.SurvivorEntity, error) {
	survivor := entitys.SurvivorEntity{Id: id, Name: input.Name, Gender: input.Gender, Lat: input.Lat, Lng: input.Lng}
	survivorFound, foundErr := s.repo.Find(id)
	if foundErr != nil {
		return entitys.SurvivorEntity{}, foundErr
	}
	newSurvivor := survivorFound.Update(survivor)
	survivorUpdated, updatedErr := s.repo.Update(newSurvivor)
	if updatedErr != nil {
		return entitys.SurvivorEntity{}, updatedErr
	}
	return survivorUpdated, nil
}
