package requestHandler

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/oyeprashar/microservice-boilerplate/structs"
	"net/http"
)

/*
	This file is used to validate the requests
*/

func HealthRequest(ctx context.Context, r *http.Request) (structs.HealthReq, error) {

	v := validator.New()
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var reqObj structs.HealthReq
	err := decoder.Decode(&reqObj)
	if err != nil {
		return reqObj, err
	}

	err = v.Struct(reqObj)
	if err != nil {
		return reqObj, err
	}
	return reqObj, nil
}
