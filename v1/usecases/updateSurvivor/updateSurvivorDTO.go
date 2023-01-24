package usecases

import (
	"serve/v1/entitys"
)

type IUpdateSurvivorController interface {
	IUpdateSurvivorCase
}

type IUpdateSurvivorCase interface {
	Execute(string, UpdateSurvivorInput) (entitys.SurvivorEntity, error)
}

type UpdateSurvivorInput struct {
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}
