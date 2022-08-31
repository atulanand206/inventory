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

func NewBuildingRouteManager(buildingConfig store.StoreConfig, routeManager *RouteManager) *BuildingRouteManager {
	return &BuildingRouteManager{
		RouteManager: *routeManager,
		service:      core.NewBuildingService(buildingConfig),
	}
}

func (rm *BuildingRouteManager) RoutesBuilding() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/buildings/init", rm.handler.postChain.Handler(rm.Create))
	router.HandleFunc("/buildings/users/add", rm.handler.postChain.Handler(rm.Add))
	router.HandleFunc("/buildings/users/remove", rm.handler.postChain.Handler(rm.Remove))
	return router
}

func (rm *BuildingRouteManager) Create(w http.ResponseWriter, r *http.Request) {
}

func (rm *BuildingRouteManager) Add(w http.ResponseWriter, r *http.Request) {
}

func (rm *BuildingRouteManager) Remove(w http.ResponseWriter, r *http.Request) {
}
