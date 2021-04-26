package handlers

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/mahadzaryab1/covid-vaccine-api/models"
)

const CANADA_POPULATION = 37057765

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

func getAllVaccinesResponse(data [][]string) models.VaccineEntries {

	response := models.VaccineEntries{VaccineData: make([]models.VaccineEntry, 0)}
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

		currEntry := models.VaccineEntry{
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
