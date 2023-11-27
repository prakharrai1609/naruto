package naruto

import (
	"net/http"

	"github.com/prakharrai1609/naruto/naruto/middlewares"
)

// App represents the naruto web application.
type App struct {
	middlewareHandlers       []middlewares.Middleware
	routeHandlers            map[string]http.Handler
	globalMiddlewareHandlers []middlewares.Middleware
	subRouters               []SubRouterWithPrefix
	uniqueID                 int
}

// SubRouterWithPrefix associates a sub-router with its prefix.
type SubRouterWithPrefix struct {
	subRouter       *App
	subRouterPrefix string
}
