package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

// JSON sets status code and returns a json formatted response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error sets status code and returns a json formatted error response
func Error(w http.ResponseWriter, statusCode int, err error) {
	val := reflect.Indirect(reflect.ValueOf(err))
	JSON(w, statusCode, struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{
		Message: err.Error(),
		Error:   val.Type().Name(),
	})
}
