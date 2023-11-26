### `App` Struct

```go
type App struct {
    middlewareHandlers []Middleware
    routeHandlers      map[string]http.Handler
}
```

- **Description:** The `App` struct represents the naruto web application.
- **Fields:**
  - `middlewareHandlers`: A slice to store middleware functions.
  - `routeHandlers`: A map to store route handlers where the key is a combination of HTTP method and path.

### `Middleware` Type

```go
type Middleware func(http.Handler) http.Handler
```

- **Description:** `Middleware` is a function type for middleware handlers. It takes an `http.Handler` and returns an `http.Handler`.

### `New` Function

```go
func New() *App {
    return &App{
        routeHandlers: make(map[string]http.Handler),
    }
}
```

- **Description:** The `New` function creates a new instance of the naruto web application.
- **Returns:** A pointer to the newly created `App` instance.

### HTTP Methods Registration Functions (`Get`, `Post`, `Put`, ...)

```go
func (app *App) Get(path string, handler http.HandlerFunc) {
    app.handle("GET", path, handler)
}

// Other HTTP methods follow the same pattern
```

- **Description:** These functions register route handlers for specific HTTP methods.
- **Parameters:**
  - `path`: The path of the route.
  - `handler`: The handler function for the route.

### `Use` Function

```go
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
```

Certainly! Below is a Markdown (MD) representation of the `UseRouter` function in your `naruto` package:

```markdown
## UseRouter Function

### Description
Registers a sub-router as middleware for a specific route.

### Signature
```go
func (app *App) UseRouter(route string, subRouter *App)
```

### Parameters
- `app`: The main application instance (receiver of the method).
- `route`: The route path where the sub-router should be mounted.
- `subRouter`: The sub-router instance to be registered as middleware.

### Usage
```go
// Create the main application instance
app := naruto.New()

// Create the sub-router instance
adminRouter := admin.NewRouter()

// Mount the sub-router under the specified route
app.UseRouter("/admin", adminRouter)
```

### Implementation
```go
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
```

### Explanation
The `UseRouter` function allows you to mount a sub-router under a specific route in the main application. It adjusts the request path before passing it to the sub-router, ensuring that the sub-router correctly handles paths relative to the mounted route.


- **Description:** The `Use` function registers a middleware for a specific route or wildcard route.
- **Parameters:**
  - `route`: The route path or a wildcard route.
  - `middleware`: The middleware function to be executed.
- **Implementation:**
  - It appends a closure to `middlewareHandlers` that checks if the requested path starts with the specified route and executes the middleware accordingly.

### `UseGlobal` Function

```go
func (app *App) UseGlobal(middleware Middleware) {
    app.middlewareHandlers = append(app.middlewareHandlers, middleware)
}
```

- **Description:** The `UseGlobal` function registers a global middleware for all routes.
- **Parameters:**
  - `middleware`: The global middleware function to be executed.

### `handle` Function

```go
func (app *App) handle(method, path string, handler http.Handler) {
    key := method + path
    app.routeHandlers[key] = handler
}
```

- **Description:** The `handle` function registers a route with the specified method, path, and handler.
- **Parameters:**
  - `method`: The HTTP method for the route.
  - `path`: The path of the route.
  - `handler`: The handler function for the route.

### `ServeHTTP` Method

```go
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
```

- **Description:** The `ServeHTTP` method implements the `http.Handler` interface for naruto `App`.
- **Parameters:**
  - `w`: The `http.ResponseWriter`.
  - `r`: The `http.Request`.
- **Implementation:**
  - It looks up the appropriate route handler based on the method and path.
  - It applies global middleware to the final handler.
  - It applies route-specific middleware in reverse order.
  - It finally serves the HTTP request using the assembled final handler.
