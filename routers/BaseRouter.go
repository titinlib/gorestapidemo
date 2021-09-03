package routers

import "github.com/gorilla/mux"

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	initializeRoutes()
}

type RouterInitializer interface {
	initialize()
}

func initializeRoutes() {
	subrouters := []RouterInitializer{
		ProductRouter{},
	}

	for _, subrsubrouter := range subrouters {
		subrsubrouter.initialize()
	}
}
