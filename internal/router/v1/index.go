package v1

import "github.com/go-chi/chi/v5"

// V1Router is router for V1 APIs
func V1Router(r chi.Router) {
	r.Route("/test", TestRouter)
}
