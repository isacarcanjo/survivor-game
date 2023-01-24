package routes

import (
	closestSurvivor "serve/v1/usecases/closestSurvivor"
	createSurvivor "serve/v1/usecases/createSurvivor"
	getSurvivor "serve/v1/usecases/getSurvivor"
	reportInfected "serve/v1/usecases/reportInfected"
	updateSurvivor "serve/v1/usecases/updateSurvivor"
)

func GetSurvivorRouters() []Router {
	var SurvivorRouters = []Router{
		{
			URI:                   "/survivor",
			Method:                "POST",
			DomainFunc:            createSurvivor.MakeCreateSurvivor().Handle,
			RequireAuthentication: true,
		},
		{
			URI:                   "/survivor/{id}",
			Method:                "GET",
			DomainFunc:            getSurvivor.MakeGetSurvivor().Handle,
			RequireAuthentication: false,
		},
		{
			URI:                   "/survivor/{id}",
			Method:                "PUT",
			DomainFunc:            updateSurvivor.MakeUpdateSurvivor().Handle,
			RequireAuthentication: false,
		},
		{
			URI:                   "/survivor/closest-survivor/{id}",
			Method:                "GET",
			DomainFunc:            closestSurvivor.MakeClosestSurvivor().Handle,
			RequireAuthentication: false,
		},
		{
			URI:                   "/survivor/infected/{id}",
			Method:                "POST",
			DomainFunc:            reportInfected.MakeReportInfected().Handle,
			RequireAuthentication: false,
		},
	}
	return SurvivorRouters
}
