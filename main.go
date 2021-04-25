package main

import (
	"encoding/json"
	"net/http"
)

type VaccineEntry struct {
	Date                       string `json:"date"`
	TotalVaccinations          int    `json:"total_vaccinations"`
	TotalPeopleVaccinated      int    `json:"total_people_vaccinated"`
	TotalPeopleFullyVaccinated int    `json:"total_people_fully_vaccinated"`
}

type VaccineEntries struct {
	VaccineData []VaccineEntry `json:"vaccine_data"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "Healthy!",
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
