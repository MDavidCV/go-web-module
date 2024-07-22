package middleware

import (
	"net/http"
	"os"

	"github.com/MDavidCV/go-web-module/internal/handler/controller"
	"github.com/MDavidCV/go-web-module/utility"
)

func AuthValidationMid(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api_key := os.Getenv("API_KEY")
		if api_key != r.Header.Get("token") || api_key == "" {
			controller.HandleResponse(w, utility.NewUnauthorizedResponse())
			return
		}

		handler.ServeHTTP(w, r)
	})
}
