# COVID-19 Vaccination API 
I built a simple REST API that provides data on Canada's COVID-19 vaccination efforts. I built this API to get familiarized with Go. The data being returned in the API was obtained from [this](https://github.com/owid/covid-19-data/blob/master/public/data/vaccinations/country_data/Canada.csv) public dataset. 

## Requirements 
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker-Compose](https://docs.docker.com/compose/install/)

## Running the API locally
The app has been dockerized and is relatively easy to get running on your local machine. Simply run the following two commands to get the API up and running on your local machine: 
```bash
docker-compose build # build the application
docker-compose up # start the application
```

To verify that the app is running as expected, make a GET request to http://localhost:8080/health. You should get the following response: 
```json
{
    "status": "Healthy!"
}
```
## Endpoints 
### GET /vaccine_data/latest
Provides a snapshot of the most recent vaccination numbers: 
```json
{
    "date": "2021-04-25",
    "total_vaccinations": 12045041,
    "total_people_vaccinated": 11026753,
    "total_people_fully_vaccinated": 1018288,
    "percentage_fully_vaccinated": 2.74784
}
```
### GET /vaccine_data
Provides all the daily available data since the start of Canada's vaccination efforts:
```json
"vaccine_data": [
    {
        "date": "2020-12-14",
        "total_vaccinations": 5,
        "total_people_vaccinated": 0,
        "total_people_fully_vaccinated": 0,
        "percentage_fully_vaccinated": 0
    }, ...
]
```

