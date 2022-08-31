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
	machineRouteManager := routes.NewMachineRouteManager(storeConfigs["machines"], routeManager)
	http.ListenAndServe(":"+os.Getenv("PORT"), machineRouteManager.RoutesMachine())
}

func collections() []string {
	return []string{"user", "buildings", "room_sharing", "building_bed", "bed_user", "machines", "usages"}
}
