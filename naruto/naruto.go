package naruto

import (
	"fmt"
	"net/http"
	"strings"
)

// App represents the naruto web application.
type App struct {
	middlewareHandlers []Middleware
	routeHandlers      map[string]http.Handler
}

// Middleware is a function type for middleware handlers.
type Middleware func(http.Handler) http.Handler

// New creates a new naruto web application.
func New() *App {
	return &App{
		routeHandlers: make(map[string]http.Handler),
	}
}

// Get registers a GET route with the specified path and handler.
func (app *App) Get(path string, handler http.HandlerFunc) {
	app.handle("GET", path, handler)
}

// Post registers a POST route with the specified path and handler.
func (app *App) Post(path string, handler http.HandlerFunc) {
	app.handle("POST", path, handler)
}

// Put registers a PUT route with the specified path and handler.
func (app *App) Put(path string, handler http.HandlerFunc) {
	app.handle("PUT", path, handler)
}

// Delete registers a DELETE route with the specified path and handler.
func (app *App) Delete(path string, handler http.HandlerFunc) {
	app.handle("DELETE", path, handler)
}

// Patch registers a PATCH route with the specified path and handler.
func (app *App) Patch(path string, handler http.HandlerFunc) {
	app.handle("PATCH", path, handler)
}

// Options registers an OPTIONS route with the specified path and handler.
func (app *App) Options(path string, handler http.HandlerFunc) {
	app.handle("OPTIONS", path, handler)
}

// Head registers a HEAD route with the specified path and handler.
func (app *App) Head(path string, handler http.HandlerFunc) {
	app.handle("HEAD", path, handler)
}

// Trace registers a TRACE route with the specified path and handler.
func (app *App) Trace(path string, handler http.HandlerFunc) {
	app.handle("TRACE", path, handler)
}

// Connect registers a CONNECT route with the specified path and handler.
func (app *App) Connect(path string, handler http.HandlerFunc) {
	app.handle("CONNECT", path, handler)
}

// Use registers a middleware for a specific route or wildcard route.
func (app *App) Use(route string, middleware Middleware) {
	app.middlewareHandlers = append(app.middlewareHandlers, func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, route) {
				middleware(next).ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	})
}

// UseRouter registers a sub-router as middleware for a specific route.
func (app *App) UseRouter(route string, subRouter *App) {
	app.middlewareHandlers = append(app.middlewareHandlers, func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.Path)
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
func (app *App) UseGlobal(middleware Middleware) {
	app.middlewareHandlers = append(app.middlewareHandlers, middleware)
}

// handle registers a route with the specified method, path, and handler.
func (app *App) handle(method, path string, handler http.Handler) {
	key := method + path
	app.routeHandlers[key] = handler
}

// ServeHTTP implements the http.Handler interface for naruto.App.
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	handler, ok := app.routeHandlers[r.Method+r.URL.Path]
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