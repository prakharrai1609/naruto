package naruto_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prakharrai1609/naruto/naruto"
)

func TestNaruto(t *testing.T) {
	app := naruto.New()

	// Register routes
	app.Get("/route1/subpath1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler for /route1/subpath1")
	})

	// Register route-level middleware 1
	// app.Use("/route1", middleware1)

	// Use middleware1 with the wrapper.
	app.Use("/route1", naruto.MiddlewareWrapper(middleware1))

	// Register route-level middleware 2
	app.Use("/route1", middleware2)

	// Register global middleware
	app.UseGlobal(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Global Middleware")
			next.ServeHTTP(w, r)
		})
	})

	// Create a test server
	server := httptest.NewServer(app)
	defer server.Close()

	// Make an HTTP request to /route1/subpath1 (updated path)
	req, err := http.NewRequest("GET", server.URL+"/route1/subpath1", nil)
	if err != nil {
		panic(err)
	}

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the response status and body
	fmt.Println("Response Status:", resp.Status)
	// Read the response body if needed
	// responseBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("Response Body:", string(responseBody))
}

// func middleware1(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Middleware 1 for /route1")
// 		// fmt.Fprintf(w, "SERVER RESPONSE MIDDLEWARE 1: %s", r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }

// middleware1 is a simple middleware function.
func middleware1(w http.ResponseWriter, r *http.Request, next http.Handler) {
	fmt.Println("Middleware 1 for /route1")
	// fmt.Fprintf(w, "SERVER RESPONSE MIDDLEWARE 1: %s", r.URL.Path)
	next.ServeHTTP(w, r)
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware 2 for /route1")
		next.ServeHTTP(w, r)
	})
}
