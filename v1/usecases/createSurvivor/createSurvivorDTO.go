package usecases

import (
	"serve/v1/entitys"
)

type ICreateSurvivorController interface {
	ICreateSurvivorCase
}

type ICreateSurvivorCase interface {
	Execute(CreateSurvivorInput) (entitys.SurvivorEntity, error)
}

type CreateSurvivorInput struct {
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}
