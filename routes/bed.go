package routes

import (
	"encoding/json"
	"net/http"

	"github.com/atulanand206/inventory/core"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type BedRouteManager struct {
	RouteManager
	service core.BedService
}

func NewBedRouteManager(
	accessConfig,
	userConfig,
	bedUserConfig store.StoreConfig,
	routeManager *RouteManager) *BedRouteManager {
	return &BedRouteManager{
		RouteManager: *routeManager,
		service:      core.NewBedService(accessConfig, userConfig, bedUserConfig),
	}
}

func (rm *BedRouteManager) RoutesBed() map[string]http.HandlerFunc {
	var routes = make(map[string]http.HandlerFunc)
	routes["/beds/create"] = rm.handler.postChain.Handler(rm.CreateAccessCode)
	routes["/beds/auth"] = rm.handler.postChain.Handler(rm.Authenticate)
	routes["/beds/users/add"] = rm.handler.postChain.Handler(rm.AddUser)
	routes["/beds/users/remove"] = rm.handler.postChain.Handler(rm.RemoveUser)
	return routes
}

func (rm *BedRouteManager) CreateAccessCode(w http.ResponseWriter, r *http.Request) {
	var bedAccess types.BedAccess
	json.NewDecoder(r.Body).Decode(&bedAccess)
	rm.service.CreateBedAccess(bedAccess)
}

func (rm *BedRouteManager) Authenticate(w http.ResponseWriter, r *http.Request) {
	var bedAccess types.BedAccess
	json.NewDecoder(r.Body).Decode(&bedAccess)
	_, err := rm.service.ValidateAccess(bedAccess)
	if err != nil {
		return
	}
	bedUser, err := rm.service.GetBedUser(bedAccess.BedId)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(bedUser)
}

func (rm *BedRouteManager) AddUser(w http.ResponseWriter, r *http.Request) {
	var addRequest types.NewAddUserRequest
	json.NewDecoder(r.Body).Decode(&addRequest)
	bedUser, err := rm.service.AddUser(addRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(bedUser)
}

func (rm *BedRouteManager) RemoveUser(w http.ResponseWriter, r *http.Request) {
	var removeRequest types.NewRemoveUserRequest
	json.NewDecoder(r.Body).Decode(&removeRequest)
	bedUser, err := rm.service.RemoveUser(removeRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(bedUser)
}
