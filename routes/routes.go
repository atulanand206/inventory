package routes

import (
	"net/http"
	"os"
	"strconv"

	net "github.com/atulanand206/go-network"
	"github.com/atulanand206/inventory/core"
	"github.com/atulanand206/inventory/role"
	"github.com/atulanand206/inventory/store"
	"github.com/atulanand206/inventory/types"
)

type RequestHandler struct {
	getChain  *net.MiddlewareChain
	postChain *net.MiddlewareChain
	putChain  *net.MiddlewareChain
}

type RouteManager struct {
	handler     *RequestHandler
	userService core.UserService
}

func New(userConfig store.StoreConfig) *RouteManager {
	// Interceptor chain for attaching to the requests.
	chain := net.MiddlewareChain{
		net.ApplicationJsonInterceptor(),
		// net.AuthenticationInterceptor(),
	}
	getChain := chain.Add(net.CorsInterceptor(http.MethodGet))
	putChain := chain.Add(net.CorsInterceptor(http.MethodPut))
	postChain := chain.Add(net.CorsInterceptor(http.MethodPost))
	routeManager := &RouteManager{
		handler: &RequestHandler{
			getChain:  &getChain,
			postChain: &postChain,
			putChain:  &putChain,
		},
		userService: core.NewUserService(store.NewUserStore(userConfig)),
	}
	return routeManager
}

func (rm *RouteManager) AssertRole(r *http.Request, capability role.Capability) bool {
	jwtClaims, err := net.Authenticate(r, rm.ClientSecret())
	if err != nil {
		return false
	}
	userRole := jwtClaims["role"].(string)
	return role.HasCapability(userRole, capability)
}

func (rm *RouteManager) CreateToken(user types.UserResponse) string {
	jwtClaims := map[string]interface{}{
		"username": user.Username,
		"role":     user.Role,
	}
	expiresIn := strconv.Itoa(24 * 60 * 60)
	token, err := net.CreateToken(jwtClaims, rm.ClientSecret(), expiresIn)
	if err != nil {
		return ""
	}
	return token
}

func (rm *RouteManager) ClientSecret() string {
	return os.Getenv("CLIENT_SECRET")
}
