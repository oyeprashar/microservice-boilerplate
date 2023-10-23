package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/oyeprashar/microservice-boilerplate/internal/controller"
	"github.com/oyeprashar/microservice-boilerplate/internal/handlers"
	"github.com/oyeprashar/microservice-boilerplate/internal/middleware/auth"
	"github.com/oyeprashar/microservice-boilerplate/internal/middleware/requestLogger"
)

// TestRouter is Router for /test pattern
func TestRouter(r chi.Router) {
	r.Route("/", testRoutes)
}

func testRoutes(r chi.Router) {

	//r.Use(requestLogger.LogRequestResponse)

	r.With(
		auth.AuthFilter,                  // This authenticates the api key,
		controller.HealthCont,            // This sets the response payload in content with key "resData"
		requestLogger.LogRequestResponse, // This will log the request and response
	).Get("/health", handlers.RespondWithSuccess)
	// This reads "resData" from context and writes to http.ResponseWriter

}
