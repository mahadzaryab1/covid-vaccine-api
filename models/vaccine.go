package models

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
