package responseHandler

import (
	"encoding/json"
	"errors"
	"github.com/oyeprashar/microservice-boilerplate/errorHandler"
	"github.com/sirupsen/logrus"
	"net/http"
	"syscall"
)

// TODO: Should be create new logger object for each file?
var log = logrus.New() // TODO: make this a singleton

// respondwithJSON : Write the response into the http.ResponseWriter object which is returned to the client
func respondwithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	defer errorHandler.Recovery(w, r, http.StatusConflict)
	response, err := json.Marshal(payload)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)

	// explore this
	if errors.Is(err, syscall.EPIPE) {
		log.Println(err)
		return
	}

	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// GenericRes : is used when we want to return the data back to client with no error
func GenericRes(w http.ResponseWriter, r *http.Request) {
	resData := r.Context().Value("resData") // read the data from the context
	var payload = map[string]interface{}{
		"status": true,
		"error":  "",
		"data":   resData,
	}

	// set the data to the ResponseWrite
	respondwithJSON(w, r, http.StatusOK, payload)
}

// RespondWithError : is used to respond back with error and no data
func RespondWithError(w http.ResponseWriter, r *http.Request, httpStatusCode int, msg string) {
	errJSON := map[string]interface{}{
		"status": false,
		"data":   map[string]interface{}{},
		"error":  msg,
	}
	respondwithJSON(w, r, httpStatusCode, errJSON)
}
