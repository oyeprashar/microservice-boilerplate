package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/oyeprashar/microservice-boilerplate/structs"
	"net/http"
)

/*
	In this file we decode the request and validate them before sending to the controllers
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
