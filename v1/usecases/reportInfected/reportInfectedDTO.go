package usecases

import (
	"serve/v1/entitys"
)

type IReportInfectedController interface {
	IReportInfectedCase
}

type IReportInfectedCase interface {
	Execute(id string) (entitys.SurvivorEntity, error)
}
