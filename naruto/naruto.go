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

// App represents the naruto web application.
type App struct {
	middlewareHandlers []middlewares.Middleware
	routeHandlers      map[string]http.Handler
	uniqueID           int
}

// New creates a new naruto web application.
func New() *App {
	idMutex.Lock()
	defer idMutex.Unlock()

	globalIDCounter++
	app := &App{
		routeHandlers: make(map[string]http.Handler),
		uniqueID:      globalIDCounter,
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

// UseRouter registers a sub-router as middleware for a specific route.
func (app *App) UseRouter(route string, subRouter *App) {
	app.middlewareHandlers = append(app.middlewareHandlers, func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, route) {
				// Adjust the request path before passing it to the sub-router
				r.URL.Path = strings.TrimPrefix(r.URL.Path, route)
				subRouter.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	})
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

/*
ServeHTTP implements the http.Handler interface for naruto.App.
It first runs the route specific middlewares.
After that it runs the global middlewares.
Finally it runs the API handler.
*/
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := RoutePathFixer(r.URL.Path)

	handler, ok := app.routeHandlers[RouteHandlerKeyGen(r.Method, path, app.uniqueID)]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Apply global middleware
	finalHandler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))

	// Apply route-specific middleware in reverse order
	for i := len(app.middlewareHandlers) - 1; i >= 0; i-- {
		finalHandler = app.middlewareHandlers[i](finalHandler)
	}

	finalHandler.ServeHTTP(w, r)
}

//new

/*
ServeHTTP implements the http.Handler interface for naruto.App.
It first runs the route specific middlewares.
After that, it runs the global middlewares.
Finally, it runs the API handler.
*/
// func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	path := RoutePathFixer(r.URL.Path)

// 	// Check for route handlers in the main app
// 	handler, ok := app.routeHandlers[RouteHandlerKeyGen(r.Method, path, app.uniqueID)]
// 	if !ok {
// 		// If not found, check sub-routers
// 		for subRouter, parentRoute := range app.subRouters {
// 			if strings.HasPrefix(r.URL.Path, parentRoute) {
// 				// Adjust the request path before passing it to the sub-router
// 				r.URL.Path = strings.TrimPrefix(r.URL.Path, parentRoute)
// 				subRouter.ServeHTTP(w, r)
// 				return
// 			}
// 		}

// 		// If no route or sub-router found, return a 404 Not Found response
// 		http.NotFound(w, r)
// 		return
// 	}

// 	// Create a stack of middleware handlers starting with the API handler
// 	middlewareStack := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		handler.ServeHTTP(w, r)
// 	}))

// 	// Apply route-specific middleware in reverse order
// 	for i := len(app.middlewareHandlers) - 1; i >= 0; i-- {
// 		middlewareStack = app.middlewareHandlers[i](middlewareStack)
// 	}

// 	// Apply global middleware in reverse order
// 	for i := len(app.globalMiddlewareHandlers) - 1; i >= 0; i-- {
// 		middlewareStack = app.globalMiddlewareHandlers[i](middlewareStack)
// 	}

// 	// Serve the request through the middleware stack
// 	middlewareStack.ServeHTTP(w, r)
// }

//new ends
