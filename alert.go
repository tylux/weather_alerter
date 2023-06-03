package main

import (
	"fmt"
	"time"
	"strings"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func sms(config Config, message string) {
	client := twilio.NewRestClient()

	//move some of these to a global config
	params := &openapi.CreateMessageParams{}
	params.SetFrom(config.twilioPhoneNumber)
	params.SetBody(message)

	sendToPhoneNumbers := strings.Split(config.sendToPhoneNumber, ",")

	for _, sendToPhoneNumber := range sendToPhoneNumbers {
		params.SetTo(sendToPhoneNumber)

		current_hour := time.Now().Hour()
		
		if current_hour >= 7 && current_hour <= 20 {
			_, err := client.Api.CreateMessage(params)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("SMS sent successfully to " + sendToPhoneNumber)
			}
		}
	}
}

