package main

import (
	"testing"
)

func TestTimeLogic(t *testing.T) {
	config := Config{
		twilioPhoneNumber: "+15017122661",
		sendToPhoneNumber: "+15558675310",
	}
	var curr_temp float64
	thresholdTemp := 70.0
	var aboveThreshold, belowThreshold bool
	var current_hour int

	// Test case 1: Current temperature is above threshold and it's morning
	curr_temp = 75.0
	current_hour = 10
	aboveThreshold = false
	belowThreshold = false
	aboveThreshold, belowThreshold = time_logic(config, curr_temp, current_hour, thresholdTemp, aboveThreshold, belowThreshold)
	if !aboveThreshold || belowThreshold {
		t.Errorf("time_logic failed, expected (true, false) but got (%t, %t)", aboveThreshold, belowThreshold)
	}

	// Test case 2: Current temperature is below threshold and it's afternoon
	curr_temp = 65.0
	current_hour = 14
	aboveThreshold = false
	belowThreshold = false
	aboveThreshold, belowThreshold = time_logic(config, curr_temp, current_hour, thresholdTemp, aboveThreshold, belowThreshold)
	if aboveThreshold || !belowThreshold {
		t.Errorf("time_logic failed, expected (false, true) but got (%t, %t)", aboveThreshold, belowThreshold)
	}

	// Test case 3: Current temperature is within threshold and it's evening
	curr_temp = 70.0
	current_hour = 20
	aboveThreshold = false
	belowThreshold = false
	aboveThreshold, belowThreshold = time_logic(config, curr_temp, current_hour, thresholdTemp, aboveThreshold, belowThreshold)
	if aboveThreshold || belowThreshold {
		t.Errorf("time_logic failed, expected (false, false) but got (%t, %t)", aboveThreshold, belowThreshold)
	}
}
