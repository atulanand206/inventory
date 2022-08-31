package routes

import (
	"encoding/json"
	"net/http"

	"github.com/atulanand206/inventory/types"
)

func (rm *RouteManager) RoutesMachine() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/machines", rm.handler.postChain.Handler(rm.GetMachines))
	router.HandleFunc("/machines/mark", rm.handler.postChain.Handler(rm.MarkMachine))
	return router
}

func (rm *RouteManager) GetMachines(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rm.dataStore.GetMachines())
}

func (rm *RouteManager) MarkMachine(w http.ResponseWriter, r *http.Request) {
	var machine types.Machine
	json.NewDecoder(r.Body).Decode(&machine)
	rm.dataStore.MarkMachine(machine)
	json.NewEncoder(w).Encode(rm.dataStore.GetMachines())
}
