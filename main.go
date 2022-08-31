package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/atulanand206/inventory/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("CORS_ORIGIN"))
	routeManager := routes.New()
	http.ListenAndServe(":"+os.Getenv("PORT"), routeManager.RoutesMachine())
}
