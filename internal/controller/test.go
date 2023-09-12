package controller

import (
	"context"
	"github.com/oyeprashar/microservice-boilerplate/errorHandler"
	"github.com/oyeprashar/microservice-boilerplate/internal/requestHandler"
	"github.com/oyeprashar/microservice-boilerplate/internal/responseHandler"
	"net/http"
)

/*
	All the controller for the /test/<> will be written here
*/

func HealthCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer errorHandler.Recovery(w, r, http.StatusConflict)

		// validate the request
		requestObj, err := requestHandler.HealthRequest(r.Context(), r)

		if err != nil {
			responseHandler.RespondWithError(w, r, http.StatusBadRequest, "bad request")
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
