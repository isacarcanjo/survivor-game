package usecases

import (
	"serve/v1/entitys"
)

type IGetSurvivorController interface {
	IGetSurvivorCase
}

type IGetSurvivorCase interface {
	Execute(id string) (entitys.SurvivorEntity, error)
}

type GetSurvivorInput struct {
	Id string `json:"id"`
}
