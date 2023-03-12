package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	open_weather_api_key string
	location string
	polling_interval int
	threshold_temp int
	send_to_phone_number string
	twilio_phone_number string
}

func main() {
	config := Config{
		polling_interval: 1, //how often to look up weather and alert
		threshold_temp: 70, //threshold at which app notifies you
		open_weather_api_key: os.Getenv("OPEN_WEATHER_API_KEY"),
		location: os.Getenv("OPEN_WEATHER_LOCATION"),
		send_to_phone_number: os.Getenv("TO_PHONE_NUMBER"),
		twilio_phone_number: os.Getenv("TWILIO_PHONE_NUMBER"),

	}

	//todo create main looper function to call twilio, openweather etc
	curr_weather := WeatherData{}
	curr_weather.getWeather(config)

	//todo write tests

}

func time_logic(config Config, curr_temp float64, threshold_temp float64, aboveThreshold bool, belowThreshold bool) (bool, bool){
	// TODO make this logic use UTC and make hour lo gic configurable
	current_time := time.Now()
    temperatureString := strconv.Itoa(int(curr_temp)) //why am I converting this
	if curr_temp > threshold_temp && int(current_time.Hour()) <= 12 && !aboveThreshold {
		message := fmt.Sprintf("Current Temp is %s°F time to close windows", temperatureString)
		fmt.Println(message)
		aboveThreshold = true
		belowThreshold = false
		sms(config, message)
	} else if curr_temp < threshold_temp && int(current_time.Hour()) >= 12 && !belowThreshold{
		message := fmt.Sprintf("Current Temp is %s°F time to open windows", temperatureString)
		fmt.Println(message)
		aboveThreshold = false
		belowThreshold = true
		sms(config, message)
	} else {
		message := fmt.Sprintf("Current Temp is %s°F nothing to do, already alerted", temperatureString)
		fmt.Println(message)		
	}
	return aboveThreshold, belowThreshold

}
