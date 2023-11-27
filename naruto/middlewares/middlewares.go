package middlewares

import "net/http"

// Middleware is a function type for middleware handlers.
type Middleware func(http.Handler) http.Handler

// MiddlewareFunc is a function type for middleware handlers.
type MiddlewareFunc func(http.ResponseWriter, *http.Request, http.Handler)

// MiddlewareWrapper wraps a function in a middleware handler.
func MiddlewareWrapper(fn MiddlewareFunc) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fn(w, r, next)
		})
	}
}
