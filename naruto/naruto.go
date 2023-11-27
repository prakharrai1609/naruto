package naruto

import (
	"net/http"
	"strings"
	"sync"

	"github.com/prakharrai1609/naruto/naruto/middlewares"
)

// globalIDCounter is a package-level variable for generating unique IDs.
var globalIDCounter int
var idMutex sync.Mutex

// New creates a new naruto web application.
func New() *App {
	idMutex.Lock()
	defer idMutex.Unlock()

	globalIDCounter++
	app := &App{
		routeHandlers:            make(map[string]http.Handler),
		globalMiddlewareHandlers: make([]middlewares.Middleware, 0),
		subRouters:               make([]SubRouterWithPrefix, 0),
		uniqueID:                 globalIDCounter,
	}
	return app
}

// Use registers a middleware for a specific route or wildcard route.
func (app *App) Use(route string, middleware middlewares.Middleware) {
	fn := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, route) {
				middleware(next).ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
	app.middlewareHandlers = append(app.middlewareHandlers, fn)
}

// UseRouter registers a sub-router in the parent router with a specific prefix.
func (app *App) UseRouter(prefix string, subRouter *App) {
	app.subRouters = append(app.subRouters, SubRouterWithPrefix{subRouter, prefix})
}

// UseGlobal registers a global middleware for all routes.
func (app *App) UseGlobal(middleware middlewares.Middleware) {
	app.middlewareHandlers = append(app.middlewareHandlers, middleware)
}

// handle registers a route with the specified method, path, and handler.
func (app *App) handle(method, path string, handler http.Handler) {
	path = RoutePathFixer(path)
	key := RouteHandlerKeyGen(method, path, app.uniqueID)
	app.routeHandlers[key] = handler
}

// ServeHTTP implements the http.Handler interface for naruto.App.
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := RoutePathFixer(r.URL.Path)

	// Check for route handlers in the main app
	handler, ok := app.routeHandlers[RouteHandlerKeyGen(r.Method, path, app.uniqueID)]
	if ok {
		// Route found in the main app, execute it
		app.executeHandler(handler, w, r)
		return
	}

	// Check for sub-routers in descending order
	for i := len(app.subRouters) - 1; i >= 0; i-- {
		subRouter := app.subRouters[i]
		if strings.HasPrefix(r.URL.Path, subRouter.subRouterPrefix) {
			// Adjust the request path before passing it to the sub-router
			r.URL.Path = strings.TrimPrefix(r.URL.Path, subRouter.subRouterPrefix)
			subRouter.subRouter.ServeHTTP(w, r)
			return
		}
	}

	// If no route or sub-router found, return a 404 Not Found response
	http.NotFound(w, r)
}

// ServeHTTP implements the http.Handler interface for naruto.App.
func (subRouter *SubRouterWithPrefix) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	subRouter.subRouter.ServeHTTP(w, r)
}

// executeHandler applies middleware and executes the final handler.
func (app *App) executeHandler(handler http.Handler, w http.ResponseWriter, r *http.Request) {
	// Create a stack of middleware handlers starting with the API handler
	middlewareStack := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))

	// Apply route-specific middleware in reverse order
	for i := len(app.middlewareHandlers) - 1; i >= 0; i-- {
		middlewareStack = app.middlewareHandlers[i](middlewareStack)
	}

	// Apply global middleware in reverse order
	for i := len(app.globalMiddlewareHandlers) - 1; i >= 0; i-- {
		middlewareStack = app.globalMiddlewareHandlers[i](middlewareStack)
	}

	// Serve the request through the middleware stack
	middlewareStack.ServeHTTP(w, r)
}
