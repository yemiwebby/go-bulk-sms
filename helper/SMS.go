package helper

import (
	"errors"
	"fmt"
	"os"
	"twilio_bulk_sms/model"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func BulkSMS(request model.SMSRequest) (string, error){

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	twilioPhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	numberOfFailedRequests := 0

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	for _, recipient := range request.Recipients{
		params := &twilioApi.CreateMessageParams{}
		params.SetTo(recipient)
		params.SetFrom(twilioPhoneNumber)
		params.SetBody(request.Message)

		_, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
			numberOfFailedRequests ++
		}
	}

	if numberOfFailedRequests > 0 {
		errorMessage := fmt.Sprintf("%d message(s) could not be sent, please check your Twilio logs for more information", numberOfFailedRequests)
		return "", errors.New(errorMessage)
	}
	
	return fmt.Sprintf("%d message(s) sent successfully", len(request.Recipients)), nil
}