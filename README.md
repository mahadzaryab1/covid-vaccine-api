# COVID-19 Vaccination API 
I built a simple REST API that provides data on Canada's COVID-19 vaccination efforts. I built this API to get familiarized with Go. The data being returned in the API was obtained from [this](https://github.com/owid/covid-19-data/blob/master/public/data/vaccinations/country_data/Canada.csv) public dataset. 

## Requirements 
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [Docker-Compose](https://docs.docker.com/compose/install/)

## Running the API locally
The app has been dockerized and is relatively easy to get running on your local machine. Simply run the following two commands to get the API up and running on your local machine: 
```bash
docker-commpose build # build the application
docker-compose up # start the application
```

To verify that the app is running as expected, make a GET request to http://localhost:8080/health. You should get the following response: 
```json
{
    "status": "Healthy!"
}
```
