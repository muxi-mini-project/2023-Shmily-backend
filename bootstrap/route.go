package bootstrap

import (
	"2023-Shmily-backend/pkg/route"
	"2023-Shmily-backend/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}
