package auth

import (
	"net/http"
)

// APIKey is the expected API key.
const APIKey = "6ecdae2c-13e2-4853-9559-4663d9f43abb" // TODO: Never hardcode secrets. Use SSM. For now we will put this in a config file

// AuthFilter is a middleware that verifies the supplied API key.
func AuthFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the API key from the request headers or query parameters, for example.
		apiKey := r.Header.Get("X-API-Key")

		// Check if the API key is valid.
		if apiKey != APIKey {
			// If the API key is invalid, return an unauthorized (401) response.
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the API key is valid, proceed to the next handler.
		next.ServeHTTP(w, r)
	})
}
