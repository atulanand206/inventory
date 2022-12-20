package routes

import (
	"encoding/json"
	"net/http"

	"github.com/atulanand206/inventory/core"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type BuildingRouteManager struct {
	RouteManager
	service core.BuildingService
}

func NewBuildingRouteManager(
	bedUserConfig,
	buildingBedConfig,
	buildingConfig,
	roomSharingConfig,
	userConfig store.StoreConfig,
	routeManager *RouteManager) *BuildingRouteManager {
	return &BuildingRouteManager{
		RouteManager: *routeManager,
		service:      core.NewBuildingService(bedUserConfig, buildingBedConfig, buildingConfig, roomSharingConfig, userConfig),
	}
}

func (rm *BuildingRouteManager) RoutesBuilding() map[string]http.HandlerFunc {
	var routes = make(map[string]http.HandlerFunc)
	routes["/buildings"] = rm.handler.postChain.Handler(rm.GetBuildings)
	routes["/buildings/init"] = rm.handler.postChain.Handler(rm.Create)
	routes["/buildings/layout"] = rm.handler.postChain.Handler(rm.GetLayout)
	routes["/buildings/users"] = rm.handler.postChain.Handler(rm.GetUsers)
	return routes
}

func (rm *BuildingRouteManager) GetBuildings(w http.ResponseWriter, r *http.Request) {
	buildings, err := rm.service.GetBuildings()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(buildings)
}

func (rm *BuildingRouteManager) Create(w http.ResponseWriter, r *http.Request) {
	var createRequest types.NewBuildingRequest
	json.NewDecoder(r.Body).Decode(&createRequest)
	building, err := rm.service.Create(createRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(building)
}

func (rm *BuildingRouteManager) GetLayout(w http.ResponseWriter, r *http.Request) {
	var request types.GetLayoutRequest
	json.NewDecoder(r.Body).Decode(&request)
	layout, err := rm.service.GetBuildingLayout(request.BuildingId)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(layout)
}

func (rm *BuildingRouteManager) GetUsers(w http.ResponseWriter, r *http.Request) {
	var request types.GetUsersForBuildingRequest
	json.NewDecoder(r.Body).Decode(&request)
	users, err := rm.service.GetUsers(request.BuildingId)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(users)
}
