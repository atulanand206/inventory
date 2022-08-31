package routes

import (
	"net/http"

	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type BuildingRouteManager struct {
	RouteManager
	dataStore store.BuildingStore
}

func NewBuildingRouteManager(config store.StoreConfig, routeManager *RouteManager) *BuildingRouteManager {
	return &BuildingRouteManager{
		RouteManager: *routeManager,
		dataStore:    store.NewMachineStoreConn(config),
	}
}

func (rm *BuildingRouteManager) RoutesBuilding() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/buildings/init", rm.handler.postChain.Handler(rm.Create))
	router.HandleFunc("/buildings/users/add", rm.handler.postChain.Handler(rm.Create))
	router.HandleFunc("/buildings/users/remove", rm.handler.postChain.Handler(rm.Create))
	return router
}

type NewBuildingRequest struct {
	Name  string      `json:"name"`
	Rooms map[int]int `json:"rooms"`
}

type NewAddUserRequest struct {
	BuildingId    string              `json:"buildingId"`
	UserId        string              `json:"userId"`
	RoomNo        int                 `json:"roomNo"`
	SharingStatus types.SharingStatus `json:"sharingStatus"`
	BedId         string              `json:"bedId"`
}

type NewRemoveUserRequest struct {
	BuildingId string `json:"buildingId"`
	UserId     string `json:"userId"`
}

func (rm *BuildingRouteManager) Create(w http.ResponseWriter, r *http.Request) {
}
