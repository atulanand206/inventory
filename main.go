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
	storeConfigs := store.StoreConfigs(os.Getenv("DB_NAME"), collections(), os.Getenv("LOCAL") == "true")
	routeManager := routes.New(storeConfigs["user"])
	router := http.NewServeMux()
	machineRouteManager := routes.NewMachineRouteManager(storeConfigs["machines"], storeConfigs["usages"], storeConfigs["bed_user"], storeConfigs["building_bed"], routeManager)
	handle(router, machineRouteManager.RoutesMachine())
	buildingRouteManager := routes.NewBuildingRouteManager(
		storeConfigs["bed_user"], storeConfigs["building_bed"], storeConfigs["buildings"], storeConfigs["room_sharing"], storeConfigs["users"], routeManager)
	handle(router, buildingRouteManager.RoutesBuilding())
	bedRouteManager := routes.NewBedRouteManager(storeConfigs["access"], storeConfigs["user"], storeConfigs["bed_user"], storeConfigs["building_bed"], routeManager)
	handle(router, bedRouteManager.RoutesBed())
	handle(router, routes.NewUserRouteManager(routeManager).RoutesUser())
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}

func handle(router *http.ServeMux, routes map[string]http.HandlerFunc) {
	for path, handler := range routes {
		router.HandleFunc(path, handler)
	}
}

func collections() []string {
	return []string{"user", "buildings", "room_sharing", "building_bed", "bed_user", "machines", "usages", "access"}
}
