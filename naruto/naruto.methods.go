package naruto

import "net/http"

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
