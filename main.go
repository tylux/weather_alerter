package main

import (
	"os"
	"strconv"
)

type Config struct {
	openWeatherApiKey string
	location          string
	pollingInterval   int
	thresholdTemp     int
	sendToPhoneNumber string
	twilioPhoneNumber string
}

func main() {
	// default vars
	pollingIntervalDefault := 5 //5 minute poll interval

    pollingInterval, _ := strconv.Atoi(os.Getenv("POLLING_INTERVAL"))
    if pollingInterval == 0 {
        pollingInterval = pollingIntervalDefault
    }

	thresholdTempDefault := 70 //temp in degrees to alert on

    thresholdTemp, _ := strconv.Atoi(os.Getenv("THRESHOLD_TEMP"))
    if thresholdTemp == 0 {
        thresholdTemp = thresholdTempDefault
    }

	config := Config{
		pollingInterval:   pollingInterval,  //how often to look up weather and alert
		thresholdTemp:     thresholdTemp, //threshold at which app notifies you
		openWeatherApiKey: os.Getenv("OPEN_WEATHER_API_KEY"),
		location:          os.Getenv("OPEN_WEATHER_LOCATION"),
		sendToPhoneNumber: os.Getenv("TO_PHONE_NUMBER"),
		twilioPhoneNumber: os.Getenv("TWILIO_PHONE_NUMBER"),
	}

	//todo create main looper function to call twilio, openweather etc
	curr_weather := WeatherData{}
	curr_weather.getWeather(config)

	//todo write tests

}
