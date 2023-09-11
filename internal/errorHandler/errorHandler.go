package errorHandler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

// TODO: Should be create new logger object for each file?
var log = logrus.New()

func writeResponse(w http.ResponseWriter, code int, message string) {

	errJSON := map[string]interface{}{
		"status": false,
		"data":   map[string]interface{}{},
		"error":  message,
	}

	response, _ := json.Marshal(errJSON)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Error(err)
	}
}

/*
	When a panic occurs in a Go program, it typically causes the program to terminate abruptly unless the panic
	is recovered. By recovering from panics, you can handle errors gracefully and prevent your program from crashing.
*/

func Recovery(w http.ResponseWriter, request *http.Request, httpStatusCode int) {
	if r := recover(); r != nil {

		// panic with message
		msg, ok := r.(string)
		if ok {
			writeResponse(w, httpStatusCode, msg)
		} else {

			// panic with error
			err, ok := r.(error)
			if ok {
				// if error object found report to sentry
				writeResponse(w, httpStatusCode, err.Error())

				// panic with something else
			} else {
				writeResponse(w, httpStatusCode, "something went wrong")
			}
		}
	}
}
