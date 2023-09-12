package errorHandler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

type respStruct struct {
	Data   map[string]interface{} `json:"data"`
	Error  interface{}            `json:"error"`
	Status bool                   `json:"status"`
}

func TestRecovery(t *testing.T) {

}

func TestWriteResponse(t *testing.T) {

	w := httptest.NewRecorder()
	writeResponse(w, 404, "resource not found")

	var resp respStruct
	err := json.Unmarshal(w.Body.Bytes(), &resp)

	if err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	if resp.Error.(string) != "resource not found" {
		t.Errorf("error expected to be : %s but got %s", "resource not found", resp.Error.(string))
	}

	if resp.Status != false {
		t.Errorf("status expected to be false but got %v", resp.Status)
	}

}
