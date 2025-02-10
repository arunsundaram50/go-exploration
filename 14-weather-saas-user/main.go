package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kelvins/geocoder"
)

type PointsResponse struct {
	Id         string
	Properties struct {
		ForecastHourly string
	}
}

type GridPointsResponse struct {
	Properties struct {
		Periods []struct {
			StartTime   string
			Temperature int
		}
	}
}

func main() {
	state := flag.String("state", "", "State for which you need the weather report")
	city := flag.String("city", "", "City for which you need the weather report")
	flag.Parse()
	fmt.Printf("State = %s, City=%s\n", *state, *city)

	geocoder.ApiKey = os.Getenv("GOOGLE_API_KEY")

	address := geocoder.Address{
		// Street:  "International Dr.",
		// Number:  1600,
		City:    *city,
		State:   *state,
		Country: "United States of America",
	}

	location, err := geocoder.Geocoding(address)
	if err != nil {
		panic(err)
	}

	fmt.Println(location.Latitude, location.Longitude)

	url := fmt.Sprintf("https://api.weather.gov/points/%f,%f", location.Latitude, location.Longitude)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var pointsResponse PointsResponse
	err = json.Unmarshal(data, &pointsResponse)
	if err != nil {
		panic(err)
	}

	gridPointsUrl := pointsResponse.Properties.ForecastHourly
	resp, err = http.Get(gridPointsUrl)
	if err != nil {
		panic(err)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var gridPointsResponse GridPointsResponse
	err = json.Unmarshal(data, &gridPointsResponse)
	if err != nil {
		panic(err)
	}

	for _, period := range gridPointsResponse.Properties.Periods {
		fmt.Println(period.StartTime, period.Temperature)
	}

}
