package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func VaccineDataHandler(w http.ResponseWriter, r *http.Request) {
	requestURL := os.Getenv("CANADA_VACCINES_URL")
	data, err := readCSV(requestURL)
	if err != nil {
		log.Print(fmt.Sprintf("ERROR: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := getAllVaccinesResponse(data)

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

func LatestVaccineDataHandler(w http.ResponseWriter, r *http.Request) {
	requestURL := os.Getenv("CANADA_VACCINES_URL")
	data, err := readCSV(requestURL)
	if err != nil {
		log.Print(fmt.Sprintf("ERROR: %s", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := getAllVaccinesResponse(data)

	jsonBytes, err := json.Marshal(response.VaccineData[len(response.VaccineData)-1])
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
