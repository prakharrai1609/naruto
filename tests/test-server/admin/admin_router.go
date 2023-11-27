package admin_test

import (
	"fmt"
	"net/http"

	"github.com/prakharrai1609/naruto/naruto"
)

func NewRouter() *naruto.App {
	adminRouter := naruto.New()

	// Define admin routes
	adminRouter.Get("/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handling GET request to /admin/create")
	})

	// Define admin routes
	adminRouter.Get("/aman", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handling GET request to /admin/aman")
	})

	return adminRouter
}
