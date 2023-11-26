package main

import (
	"fmt"
	"net/http"

	"github.com/prakharrai1609/naruto/naruto"
	admin_test "github.com/prakharrai1609/naruto/tests/test-server/admin"
)

func main() {
	app := naruto.New()

	// Create admin router
	adminRouter := admin_test.NewRouter()

	// Mount the admin router under /admin
	app.UseRouter("/admin", adminRouter)

	app.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", "Naruto app is healthy")
	})

	// Start the server
	server := http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	fmt.Println("Server listening on :8080")
	server.ListenAndServe()
}
