package routes

import (
	"net/http"

	"github.com/atulanand206/inventory/core"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
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

func (rm *BuildingRouteManager) Add(w http.ResponseWriter, r *http.Request) {
}

func (rm *BuildingRouteManager) Remove(w http.ResponseWriter, r *http.Request) {
}
