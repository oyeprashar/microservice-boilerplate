package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/oyeprashar/microservice-boilerplate/internal/controller"
	"github.com/oyeprashar/microservice-boilerplate/internal/middleware/auth"
	"github.com/oyeprashar/microservice-boilerplate/internal/responseHandler"
)

// TestRouter is Router for /test pattern
func TestRouter(r chi.Router) {
	r.Route("/", testRoutes)
}

func testRoutes(r chi.Router) {

	r.With(
		auth.AuthFilter,       // This authenticates the api key,
		controller.HealthCont, // This sets the response payload in content with key "resData"
	).Get("/health", responseHandler.GenericRes)
	// This reads "resData" from context and writes to http.ResponseWriter

}
