package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "Healthy!",
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		log.Print(fmt.Sprintf("ERROR: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
