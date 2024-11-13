package resp

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("Error with encode Response: ", err)
		return
	}
}
