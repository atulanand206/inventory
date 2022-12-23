package routes

import (
	"encoding/json"
	"net/http"

	"github.com/atulanand206/inventory/core"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type UserRouteManager struct {
	RouteManager
	service core.UserService
}

func NewUserRouteManager(userConfig store.StoreConfig, routeManager *RouteManager) *UserRouteManager {
	return &UserRouteManager{
		RouteManager: *routeManager,
		service:      core.NewUserService(store.NewUserStore(userConfig)),
	}
}

func (rm *UserRouteManager) RoutesUser() map[string]http.HandlerFunc {
	var routes = make(map[string]http.HandlerFunc)
	routes["/users/init"] = rm.handler.postChain.Handler(rm.Create)
	routes["/users/one"] = rm.handler.postChain.Handler(rm.GetUser)
	routes["/users"] = rm.handler.postChain.Handler(rm.GetUsers)
	routes["/users/login"] = rm.handler.postChain.Handler(rm.Login)
	routes["/users/reset"] = rm.handler.postChain.Handler(rm.Reset)
	return routes
}

func (rm *UserRouteManager) Create(w http.ResponseWriter, r *http.Request) {
	var createRequest types.CreateUserRequest
	json.NewDecoder(r.Body).Decode(&createRequest)
	user, err := rm.service.CreateUser(createRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (rm *UserRouteManager) GetUser(w http.ResponseWriter, r *http.Request) {
	var request types.GetUserRequest
	json.NewDecoder(r.Body).Decode(&request)
	user, err := rm.service.GetUser(request.Username)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (rm *UserRouteManager) GetUsers(w http.ResponseWriter, r *http.Request) {
	var request types.GetUsersRequest
	json.NewDecoder(r.Body).Decode(&request)
	users, err := rm.service.GetUsers(request.Usernames)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (rm *UserRouteManager) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest types.LoginRequest
	json.NewDecoder(r.Body).Decode(&loginRequest)
	user, err := rm.service.LoginUser(loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (rm *UserRouteManager) Reset(w http.ResponseWriter, r *http.Request) {
	var resetRequest types.ResetPasswordRequest
	json.NewDecoder(r.Body).Decode(&resetRequest)
	err := rm.service.ResetPassword(resetRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode("Password reset successful")
}
