package usecases

import (
	"serve/v1/entitys"
)

type IClosestSurvivorController interface {
	IClosestSurvivorCase
}

type IClosestSurvivorCase interface {
	Execute(id string) (entitys.SurvivorEntity, error)
}

type ClosestSurvivorInput struct {
	Id string `json:"id"`
}
