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

func (rm *MachineRouteManager) RoutesMachine() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/machines/init", rm.handler.postChain.Handler(rm.CreateMachines))
	router.HandleFunc("/machines", rm.handler.postChain.Handler(rm.GetMachines))
	router.HandleFunc("/machines/mark", rm.handler.postChain.Handler(rm.MarkMachine))
	return router
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

type MarkMachineRequest struct {
	MachineId string       `json:"machineId"`
	UserId    string       `json:"userId"`
	Status    types.Status `json:"status"`
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
