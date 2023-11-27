package user_test

import (
	"fmt"
	"net/http"

	"github.com/prakharrai1609/naruto/naruto"
	sub_user_test "github.com/prakharrai1609/naruto/tests/test-server/user/subUser"
)

func NewRouter() *naruto.App {
	userRouter := naruto.New()

	userRouter.Get("", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handling GET request to /user")
	})

	// Define admin routes
	userRouter.Get("/class", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Handling GET request to /user/class")
	})

	subUserRouter := sub_user_test.SubUserRouter()
	userRouter.UseRouter("/test", subUserRouter)

	return userRouter
}
