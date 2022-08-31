package routes

import (
	"net/http"

	"github.com/atulanand206/inventory/core"
	"github.com/atulanand206/inventory/store"
)

type BuildingRouteManager struct {
	RouteManager
	service core.BuildingService
}

func NewBuildingRouteManager(buildingConfig store.StoreConfig, userConfig store.StoreConfig, routeManager *RouteManager) *BuildingRouteManager {
	return &BuildingRouteManager{
		RouteManager: *routeManager,
		service:      core.NewBuildingService(buildingConfig, userConfig),
	}
}

func (rm *BuildingRouteManager) RoutesBuilding() map[string]http.HandlerFunc {
	var routes = make(map[string]http.HandlerFunc)
	routes["/buildings/init"] = rm.handler.postChain.Handler(rm.Create)
	routes["/buildings/users/add"] = rm.handler.postChain.Handler(rm.Add)
	routes["/buildings/users/remove"] = rm.handler.postChain.Handler(rm.Remove)
	return routes
}

func (rm *BuildingRouteManager) Create(w http.ResponseWriter, r *http.Request) {
}

func (rm *BuildingRouteManager) Add(w http.ResponseWriter, r *http.Request) {
}

func (rm *BuildingRouteManager) Remove(w http.ResponseWriter, r *http.Request) {
}
