package controller

import (
	"context"
	"github.com/oyeprashar/microservice-boilerplate/errorHandler"
	"github.com/oyeprashar/microservice-boilerplate/internal/handlers"
	"net/http"
)

/*
	All the controller for the /test/<endpoint> will be written here
*/

// HealthCont : is a test controller
func HealthCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorHandler.Recovery(w, r, http.StatusConflict)

		// validate the request
		requestObj, err := handlers.HealthRequest(r.Context(), r)

		if err != nil {
			handlers.RespondWithError(w, r, http.StatusBadRequest, "bad request")
			return
		}

		// TODO: remove this code
		//time.Sleep(time.Second * 10)

		// If the request was successfully validated we can proceed with required process
		var resData = map[string]string{
			"received_name": requestObj.Name,
		}

		// Setting the resData into context
		ctx := context.WithValue(r.Context(), "resData", resData)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
