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

func NewMachineRouteManager(machineConfig store.StoreConfig, routeManager *RouteManager) *MachineRouteManager {
	return &MachineRouteManager{
		RouteManager: *routeManager,
		service:      core.NewMachineService(machineConfig),
	}
}

func (rm *MachineRouteManager) RoutesMachine() map[string]http.HandlerFunc {
	var routes = make(map[string]http.HandlerFunc)
	routes["/machines/init"] = rm.handler.postChain.Handler(rm.CreateMachines)
	routes["/machines"] = rm.handler.postChain.Handler(rm.GetMachines)
	routes["/machines/mark"] = rm.handler.postChain.Handler(rm.MarkMachine)
	return routes
}

func (rm *MachineRouteManager) CreateMachines(w http.ResponseWriter, r *http.Request) {
	var machines []types.Machine
	json.NewDecoder(r.Body).Decode(&machines)
	rm.service.CreateMachines(machines)
	machines, err := rm.service.GetMachines()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(machines)
}

func (rm *MachineRouteManager) GetMachines(w http.ResponseWriter, r *http.Request) {
	machines, err := rm.service.GetMachines()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(machines)
}

func (rm *MachineRouteManager) MarkMachine(w http.ResponseWriter, r *http.Request) {
	var machine types.Machine
	json.NewDecoder(r.Body).Decode(&machine)
	rm.service.MarkMachine(machine)
	machines, err := rm.service.GetMachines()
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(machines)
}
