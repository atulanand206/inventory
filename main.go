package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	net "github.com/atulanand206/go-network"
	"github.com/joho/godotenv"
)

type Status int

const (
	Free Status = iota
	Busy
	OutOfService
)

func (s Status) String() string {
	switch s {
	case Free:
		return "Free"
	case Busy:
		return "Busy"
	case OutOfService:
		return "OutOfService"
	}
	return "unknown"
}

type Machine struct {
	Name   string `json:"name"`
	No     int    `json:"id"`
	Status Status `json:"status"`
}

type Inventory struct {
	Machines []Machine
}

var Machines = []Machine{
	{Name: "Striker", No: 1, Status: Free},
	{Name: "Maverick", No: 2, Status: Free},
	{Name: "Diamond", No: 3, Status: Free},
	{Name: "Highland", No: 4, Status: Free},
}

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("CORS_ORIGIN"))
	// Interceptor chain for attaching to the requests.
	chain := net.MiddlewareChain{
		net.ApplicationJsonInterceptor(),
		// net.AuthenticationInterceptor(),
	}

	// Interceptor chain with only PUT method.
	// getChain := chain.Add(net.CorsInterceptor(http.MethodGet))
	// putChain := chain.Add(net.CorsInterceptor(http.MethodPut))
	// Interceptor chain with only POST method.
	postChain := chain.Add(net.CorsInterceptor(http.MethodPost))
	router := http.NewServeMux()
	router.HandleFunc("/machines", postChain.Handler(getMachines))
	router.HandleFunc("/machines/mark", postChain.Handler(markMachine))
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}

func getMachines(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Machines)
}

func markMachine(w http.ResponseWriter, r *http.Request) {
	var machine Machine
	json.NewDecoder(r.Body).Decode(&machine)
	for i, m := range Machines {
		if m.No == machine.No {
			Machines[i].Status = machine.Status
		}
	}
	json.NewEncoder(w).Encode(Machines)
}
