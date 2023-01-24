package usecases

import (
	"serve/v1/entitys"
	"serve/v1/repositories"
)

type ClosestUserCase struct {
	repo repositories.ISurvivorRepository
}

func NewService(r repositories.ISurvivorRepository) *ClosestUserCase {
	return &ClosestUserCase{
		repo: r,
	}
}

func (s *ClosestUserCase) Execute(id string) (entitys.SurvivorEntity, error) {
	survivor, getErr := s.repo.Find(id)
	if getErr != nil {
		return entitys.SurvivorEntity{}, getErr
	}
	survivors, findAllErr := s.repo.FindAll()

	if findAllErr != nil {
		return entitys.SurvivorEntity{}, findAllErr
	}
	var closestSurvivor entitys.SurvivorEntity
	var lastDist float64
	for _, service := range survivors {
		if survivor.Id == service.Id {
			continue
		}
		dist := survivor.Distance(service.Lat, service.Lng, "K")
		// fmt.Printf("%s, %#v\n", service.Name, dist)
		if lastDist == 0 || dist < lastDist {
			lastDist = dist
			closestSurvivor = service
		}
	}
	return closestSurvivor, nil
}
