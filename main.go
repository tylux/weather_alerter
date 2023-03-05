package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	apiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	location := os.Getenv("OPEN_WEATHER_LOCATION")
	refresh_time := 1    //minutes
	threshold_temp := 50 //farenheit

	//todo create main looper function to call twilio, openweather etc
	curr_weather := WeatherData{}
	curr_weather.getWeather(apiKey, location, refresh_time, threshold_temp)

}

func time_logic(curr_temp float64, threshold_temp float64, aboveThreshold bool, belowThreshold bool) (bool, bool){
	// TODO make this logic use UTC and make hour logic configurable
	current_time := time.Now()
    temperatureString := strconv.Itoa(int(curr_temp))

	if curr_temp > threshold_temp && int(current_time.Hour()) < 12 && !aboveThreshold {
		message := fmt.Sprintf("Current Temp is %s°F time to close windows", temperatureString)
		fmt.Println(message)
		aboveThreshold = true
		belowThreshold = false
	} else if curr_temp < threshold_temp && int(current_time.Hour()) > 12 && !belowThreshold{
		message := fmt.Sprintf("Current Temp is %s°F time to open windows", temperatureString)
		fmt.Println(message)
		aboveThreshold = false
		belowThreshold = true
	} else {
		message := fmt.Sprintf("Current Temp is %s°F nothing to do, already alerted", temperatureString)
		fmt.Println(message)
		//sms(message)
		
	}
	return aboveThreshold, belowThreshold

}
