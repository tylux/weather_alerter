package main

import (
	"fmt"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func sms(config Config, message string) {
	client := twilio.NewRestClient()

	//move some of these to a global config
	params := &openapi.CreateMessageParams{}
	params.SetTo(config.sendToPhoneNumber)
	params.SetFrom(config.twilioPhoneNumber)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
