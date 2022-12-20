package main

import (
	"net/http"
	"os"

	"github.com/atulanand206/inventory/routes"
	"github.com/atulanand206/inventory/store"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routeManager := routes.New()
	storeConfigs := store.StoreConfigs(os.Getenv("DB_NAME"), collections(), os.Getenv("LOCAL") == "true")
	router := http.NewServeMux()
	machineRouteManager := routes.NewMachineRouteManager(storeConfigs["machines"], storeConfigs["usages"], routeManager)
	buildingRouteManager := routes.NewBuildingRouteManager(
		storeConfigs["bed_user"], storeConfigs["building_bed"], storeConfigs["buildings"], storeConfigs["room_sharing"], storeConfigs["users"], routeManager)
	handle(router, machineRouteManager.RoutesMachine())
	handle(router, buildingRouteManager.RoutesBuilding())
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}

func handle(router *http.ServeMux, routes map[string]http.HandlerFunc) {
	for path, handler := range routes {
		router.HandleFunc(path, handler)
	}
}

func collections() []string {
	return []string{"user", "buildings", "room_sharing", "building_bed", "bed_user", "machines", "usages"}
}
