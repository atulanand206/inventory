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

func NewBuildingRouteManager(buildingConfig store.StoreConfig, userConfig store.StoreConfig, routeManager *RouteManager) *BuildingRouteManager {
	return &BuildingRouteManager{
		RouteManager: *routeManager,
		service:      core.NewBuildingService(buildingConfig, userConfig),
	}
}

func (rm *BuildingRouteManager) RoutesBuilding() map[string]http.HandlerFunc {
	var routes = make(map[string]http.HandlerFunc)
	routes["/buildings/init"] = rm.handler.postChain.Handler(rm.Create)
	routes["/buildings/users"] = rm.handler.postChain.Handler(rm.GetUsers)
	routes["/buildings/users/add"] = rm.handler.postChain.Handler(rm.AddUser)
	routes["/buildings/users/remove"] = rm.handler.postChain.Handler(rm.RemoveUser)
	return routes
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

func (rm *BuildingRouteManager) GetUsers(w http.ResponseWriter, r *http.Request) {
	var request types.GetUsersForBuildingRequest
	json.NewDecoder(r.Body).Decode(&request)
	users, err := rm.service.GetUsers(request.BuildingId)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (rm *BuildingRouteManager) AddUser(w http.ResponseWriter, r *http.Request) {
	var addRequest types.NewAddUserRequest
	json.NewDecoder(r.Body).Decode(&addRequest)
	bedUser, err := rm.service.AddUser(addRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(bedUser)
}

func (rm *BuildingRouteManager) RemoveUser(w http.ResponseWriter, r *http.Request) {
	var removeRequest types.NewRemoveUserRequest
	json.NewDecoder(r.Body).Decode(&removeRequest)
	bedUser, err := rm.service.RemoveUser(removeRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(bedUser)
}
