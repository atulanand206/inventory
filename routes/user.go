package routes

import (
	"encoding/json"
	"net/http"

	"github.com/atulanand206/inventory/role"
	"github.com/atulanand206/inventory/types"
)

type UserRouteManager struct {
	RouteManager
}

func NewUserRouteManager(routeManager *RouteManager) *UserRouteManager {
	return &UserRouteManager{
		RouteManager: *routeManager,
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
	rm.AssertRole(r, role.User_Create)
	var createRequest types.CreateUserRequest
	json.NewDecoder(r.Body).Decode(&createRequest)
	user, err := rm.userService.CreateUser(createRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (rm *UserRouteManager) GetUser(w http.ResponseWriter, r *http.Request) {
	rm.AssertRole(r, role.User_Get)
	var request types.GetUserRequest
	json.NewDecoder(r.Body).Decode(&request)
	user, err := rm.userService.GetUser(request.Username)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (rm *UserRouteManager) GetUsers(w http.ResponseWriter, r *http.Request) {
	rm.AssertRole(r, role.User_Get)
	var request types.GetUsersRequest
	json.NewDecoder(r.Body).Decode(&request)
	users, err := rm.userService.GetUsers(request.Usernames)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (rm *UserRouteManager) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest types.LoginRequest
	json.NewDecoder(r.Body).Decode(&loginRequest)
	user, err := rm.userService.LoginUser(loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	token := rm.CreateToken(user)
	w.Header().Set("Authorization", token)
	json.NewEncoder(w).Encode(user)
}

func (rm *UserRouteManager) Reset(w http.ResponseWriter, r *http.Request) {
	var resetRequest types.ResetPasswordRequest
	json.NewDecoder(r.Body).Decode(&resetRequest)
	err := rm.userService.ResetPassword(resetRequest)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode("Password reset successful")
}
