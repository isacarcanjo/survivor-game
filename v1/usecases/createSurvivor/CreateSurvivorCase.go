package usecases

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type CreateUserCase struct {
	repo repositories.ISurvivorRepository
}

func NewService(r repositories.ISurvivorRepository) *CreateUserCase {
	return &CreateUserCase{
		repo: r,
	}
}

func (s *CreateUserCase) Execute(input CreateSurvivorInput) (entitys.SurvivorEntity, error) {
	survivor, err := entitys.NewSurvivor(input.Name, input.Gender, input.Lat, input.Lng)
	if err != nil {
		return survivor, err
	}
	newSurvivor, createErr := s.repo.Create(survivor)
	if createErr != nil {
		return entitys.SurvivorEntity{}, createErr
	}
	return newSurvivor, createErr
}
