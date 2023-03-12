package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (w *WeatherData) getWeather(config Config) {
	ticker := time.NewTicker(time.Duration(config.polling_interval) * time.Minute)

	// used to only alert once as the temp goes above an then again below the desired temp
	var aboveThreshold, belowThreshold bool

	fmt.Printf("Getting Weather data for %s on time interval: %d minutes\n", config.location, config.polling_interval)
	for range ticker.C {
	
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", config.location, config.open_weather_api_key)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var weatherData WeatherData
		if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
			panic(err)
		}
		temperatureF := (weatherData.Main.Temp-273.15)*1.8 + 32

		aboveThreshold, belowThreshold = time_logic(config, temperatureF, float64(config.threshold_temp), aboveThreshold, belowThreshold)
	}
}

type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
