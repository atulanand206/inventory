package routes

import (
	"net/http"

	net "github.com/atulanand206/go-network"
)

type RequestHandler struct {
	getChain  *net.MiddlewareChain
	postChain *net.MiddlewareChain
	putChain  *net.MiddlewareChain
}

type RouteManager struct {
	handler *RequestHandler
}

func New() *RouteManager {
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
	}
	return routeManager
}
