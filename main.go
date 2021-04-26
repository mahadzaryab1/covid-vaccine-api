package main

import (
	"net/http"
	"os"

	"github.com/mahadzaryab1/covid-vaccine-api/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/vaccine_data", handlers.VaccineDataHandler)
	http.HandleFunc("/vaccine_data/latest", handlers.LatestVaccineDataHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
