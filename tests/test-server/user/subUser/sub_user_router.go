package sub_user_test

import (
	"fmt"
	"net/http"

	"github.com/prakharrai1609/naruto/naruto"
)

func SubUserRouter() *naruto.App {
	subUserRouter := naruto.New()

	// Define admin routes
	subUserRouter.Get("/final", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handling GET request to /user/test/final")
	})

	return subUserRouter
}
