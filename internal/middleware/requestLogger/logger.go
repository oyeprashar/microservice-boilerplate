package requestLogger

import (
	"io"
	"log"
	"net/http"
)

/*
	The problem is that when we read from the buffer, the buffer gets empty and the controller gets empty request
	which will lead to a bad request
*/

/*
	TODO: We need to write the following components
			1. WithRequest: This will be used by the handler that will log the request when it handles it
*/

func LogRequestResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details

		//logrusLogger := logrus.New()
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		log.Println("----> checking the request body from the buffer =", string(body))

		//marshalledBody, _ := json.Marshal(r.Body.Read())
		//
		//logrusLogger.Info("Request:", r.Method, r.URL.Path, string(body))
		//
		//// Create a response recorder to capture the response
		//recorder := httptest.NewRecorder()
		//
		//// Call the next handler in the chain and pass the recorder instead of the original writer
		//next.ServeHTTP(recorder, r)
		//
		//// Log the response details
		//logrusLogger.Info("Response:", recorder.Code)
		//
		//// Copy the recorded response back to the original writer
		//for k, v := range recorder.Header() {
		//	w.Header()[k] = v
		//}
		//
		//r.Body = io.NopCloser(bytes.NewBuffer(body))
		//w.WriteHeader(recorder.Code)
		//w.Write(recorder.Body.Bytes())
	})
}
