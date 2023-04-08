package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (w *WeatherData) getWeather(config Config) {
	ticker := time.NewTicker(time.Duration(config.pollingInterval) * time.Minute)

	// used to only alert once as the temp goes above an then again below the desired temp
	var aboveThreshold, belowThreshold bool

	fmt.Printf("Getting Weather data for %s on time interval: %d minutes\n", config.location, config.pollingInterval)
	for range ticker.C {

		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", config.location, config.openWeatherApiKey)

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
		//current_time := time.Now()
		current_hour := time.Now().Hour()
		aboveThreshold, belowThreshold = time_logic(config, temperatureF, current_hour, float64(config.thresholdTemp), aboveThreshold, belowThreshold)
	}
}

func time_logic(config Config, curr_temp float64, current_hour int, thresholdTemp float64, aboveThreshold bool, belowThreshold bool) (bool, bool) {
	// TODO make this logic use UTC and make hour lo gic configurable
	temperatureString := strconv.Itoa(int(curr_temp)) //why am I converting this
	// morning alert to close windows
	if curr_temp > thresholdTemp && current_hour <= 12 && !aboveThreshold {
		message := fmt.Sprintf("Current Temp is %s°F time to close windows", temperatureString)
		fmt.Println(message)
		aboveThreshold = true
		belowThreshold = false
		sms(config, message)
	} else if curr_temp < thresholdTemp && current_hour >= 12 && !belowThreshold {
		message := fmt.Sprintf("Current Temp is %s°F time to open windows", temperatureString)
		fmt.Println(message)
		aboveThreshold = false
		belowThreshold = true
		sms(config, message)
	} else {
		message := fmt.Sprintf("Time %d, Current Temp is %s°F, threshold is %f,  nothing to do, already alerted", current_hour, temperatureString, thresholdTemp)
		fmt.Println(message)
	}
	return aboveThreshold, belowThreshold

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
