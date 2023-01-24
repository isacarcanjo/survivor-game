package usecases

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type ReportInfectedCase struct {
	repo repositories.ISurvivorRepository
}

func NewService(r repositories.ISurvivorRepository) *ReportInfectedCase {
	return &ReportInfectedCase{
		repo: r,
	}
}

func (s *ReportInfectedCase) Execute(id string) (entitys.SurvivorEntity, error) {
	survivor, getErr := s.repo.Find(id)
	if getErr != nil {
		return entitys.SurvivorEntity{}, getErr
	}
	newSurvivor := survivor.ReportInfected()
	survivorUpdated, updateErr := s.repo.Update(newSurvivor)
	if updateErr != nil {
		return entitys.SurvivorEntity{}, updateErr
	}
	return survivorUpdated, nil
}
