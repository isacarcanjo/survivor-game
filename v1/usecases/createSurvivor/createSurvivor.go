package usecases

import (
	config2 "serve/v1/config"
	"serve/v1/providers/implementations"
	repositories "serve/v1/repositories/implements"
)

func MakeCreateSurvivor() CreateSurvivorController {
	config := config2.GetDotEnv()
	db := implementations.SimDBProvider{SurvivorCollection: config.SurvivorCollection}
	userRepo := repositories.NewSurvivorRepository(db)
	makeCreateSurvivorCase := NewService(userRepo)
	makeController := CreateSurvivorController{ICreateSurvivorController: makeCreateSurvivorCase}
	return makeController
}
