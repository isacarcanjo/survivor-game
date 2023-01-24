package usecases

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type GetUserCase struct {
	repo repositories.ISurvivorRepository
}

func NewService(r repositories.ISurvivorRepository) *GetUserCase {
	return &GetUserCase{
		repo: r,
	}
}

func (s *GetUserCase) Execute(id string) (entitys.SurvivorEntity, error) {
	survivor, getErr := s.repo.Find(id)
	if getErr != nil {
		return entitys.SurvivorEntity{}, getErr
	}
	return survivor, nil
}
