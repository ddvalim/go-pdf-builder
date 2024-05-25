package response

import (
	"encoding/json"
	"github.com/ddvalim/go-pdf-builder/core/ports"
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, data interface{}) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
	}

	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Error(w http.ResponseWriter, message string, statusCode int) {
	var response ports.Response

	response.Message = message
	response.StatusCode = statusCode

	Write(w, statusCode, response)
}
