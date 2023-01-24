package usecases

import (
	config2 "serve/v1/config"
	"serve/v1/providers/implementations"
	repositories "serve/v1/repositories/implements"
)

func MakeReportInfected() ReportInfectedController {
	config := config2.GetDotEnv()
	db := implementations.SimDBProvider{SurvivorCollection: config.SurvivorCollection}
	userRepo := repositories.NewSurvivorRepository(db)
	makeReportInfectedCase := NewService(userRepo)
	makeController := ReportInfectedController{IReportInfectedController: makeReportInfectedCase}
	return makeController
}
