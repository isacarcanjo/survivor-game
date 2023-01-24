package routes

func GetRoutes() []Router {
	var routes = append([]Router{}, GetSurvivorRouters()[:]...)
	return routes
}

type Routes struct {
	Items []Router
}
