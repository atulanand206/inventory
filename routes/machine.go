package routes

import (
	"encoding/json"
	"net/http"

	"github.com/atulanand206/inventory/core"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type MachineRouteManager struct {
	RouteManager
	service core.MachineService
}

func NewMachineRouteManager(machineConfig, usageConfig, bedUserConfig, buildingBedConfig store.StoreConfig, routeManager *RouteManager) *MachineRouteManager {
	return &MachineRouteManager{
		RouteManager: *routeManager,
		service:      core.NewMachineService(machineConfig, usageConfig, bedUserConfig, buildingBedConfig),
	}
}

func (rm *MachineRouteManager) RoutesMachine() map[string]http.HandlerFunc {
	routes := make(map[string]http.HandlerFunc)
	routes["/machines/init"] = rm.handler.postChain.Handler(rm.CreateMachines)
	routes["/machines"] = rm.handler.postChain.Handler(rm.GetMachines)
	routes["/machines/mark"] = rm.handler.postChain.Handler(rm.MarkMachine)
	routes["/machines/unmark"] = rm.handler.postChain.Handler(rm.UnMarkMachine)
	return routes
}

func (rm *MachineRouteManager) CreateMachines(w http.ResponseWriter, r *http.Request) {
	var request types.CreateMachinesRequest
	json.NewDecoder(r.Body).Decode(&request)
	rm.service.CreateMachines(request)
	machines, err := rm.service.GetMachines(request.BuildingId)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(machines)
}

func (rm *MachineRouteManager) GetMachines(w http.ResponseWriter, r *http.Request) {
	var request types.GetMachinesRequest
	json.NewDecoder(r.Body).Decode(&request)
	machines, err := rm.service.GetMachines(request.BuildingId)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(machines)
}

func (rm *MachineRouteManager) MarkMachine(w http.ResponseWriter, r *http.Request) {
	var machine types.MarkMachineRequest
	json.NewDecoder(r.Body).Decode(&machine)
	rm.service.MarkMachine(machine)
	machines, err := rm.service.GetMachines(machine.BuildingId)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(machines)
}

func (rm *MachineRouteManager) UnMarkMachine(w http.ResponseWriter, r *http.Request) {
	var machine types.MarkMachineRequest
	json.NewDecoder(r.Body).Decode(&machine)
	rm.service.UnMarkMachine(machine)
	machines, err := rm.service.GetMachines(machine.BuildingId)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(machines)
}
