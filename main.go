package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

const CANADA_POPULATION = 37057765

type VaccineEntry struct {
	Date                       string  `json:"date"`
	TotalVaccinations          int     `json:"total_vaccinations"`
	TotalPeopleVaccinated      int     `json:"total_people_vaccinated"`
	TotalPeopleFullyVaccinated int     `json:"total_people_fully_vaccinated"`
	FullyVaccinatedPercentage  float32 `json:"percentage_fully_vaccinated"`
}

type VaccineEntries struct {
	VaccineData []VaccineEntry `json:"vaccine_data"`
}

func readCSV(csvUrl string) ([][]string, error) {
	response, err := http.Get(csvUrl)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	reader := csv.NewReader(response.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func getAllVaccinesResponse(data [][]string) VaccineEntries {

	response := VaccineEntries{VaccineData: make([]VaccineEntry, 0)}
	for idx, row := range data {
		if idx == 0 {
			continue
		}

		totalVaccinations, _ := strconv.Atoi(row[4])
		totalPeopleVaccinated, _ := strconv.Atoi(row[5])
		totalPeopleFullyVaccinated, _ := strconv.Atoi(row[6])
		fullyVaccinatedPercentage := float32(0)
		if totalPeopleVaccinated > 0 {
			fullyVaccinatedPercentage = (float32(totalPeopleFullyVaccinated) / float32(CANADA_POPULATION)) * 100
		}

		currEntry := VaccineEntry{
			Date:                       row[1],
			TotalVaccinations:          totalVaccinations,
			TotalPeopleVaccinated:      totalPeopleVaccinated,
			TotalPeopleFullyVaccinated: totalPeopleFullyVaccinated,
			FullyVaccinatedPercentage:  fullyVaccinatedPercentage,
		}

		response.VaccineData = append(response.VaccineData, currEntry)
	}

	return response
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

func vaccineDataHandler(w http.ResponseWriter, r *http.Request) {
	requestURL := os.Getenv("CANADA_VACCINES_URL")
	data, err := readCSV(requestURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := getAllVaccinesResponse(data)

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

func latestVaccineDataHandler(w http.ResponseWriter, r *http.Request) {
	requestURL := os.Getenv("CANADA_VACCINES_URL")
	data, err := readCSV(requestURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := getAllVaccinesResponse(data)

	jsonBytes, err := json.Marshal(response.VaccineData[len(response.VaccineData)-1])
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
	http.HandleFunc("/vaccine_data", vaccineDataHandler)
	http.HandleFunc("/vaccine_data/latest", latestVaccineDataHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
