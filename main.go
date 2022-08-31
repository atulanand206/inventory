package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/atulanand206/inventory/routes"
	"github.com/atulanand206/inventory/store"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("CORS_ORIGIN"))
	routeManager := routes.New()
	machineRouteManager := routes.NewRM(store.StoreConfig{
		DbName:    os.Getenv("DB_NAME"),
		TableName: os.Getenv("TABLE_NAME_MACHINE"),
		Local:     os.Getenv("LOCAL") == "true",
	}, routeManager)
	http.ListenAndServe(":"+os.Getenv("PORT"), machineRouteManager.RoutesMachine())
}
