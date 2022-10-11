package model

import (
	"errors"
	"log"
	"regexp"
)

type SMSRequest struct {
	Recipients []string `json:"recipients"`
	Message string `json:"message"`
}

func (smsRequest *SMSRequest) Validate() error {
	if smsRequest.Message == "" {
		return errors.New("message cannot be empty")
	}
	err := validatePhoneNumbers(smsRequest.Recipients)
	if err != nil{
		return err
	}
	return nil
}

func validatePhoneNumbers (phoneNumbers []string) error {
	for _,phoneNumber := range(phoneNumbers){
		err := validatePhoneNumber(phoneNumber)
		if err != nil {
			return errors.New("all phone numbers must be in E.164 format")
		}
	}
	return nil
}

func validatePhoneNumber(phoneNumber string) error{
	e164Pattern := `^\+[1-9]\d{1,14}$`
	match, err := regexp.Match(e164Pattern, []byte(phoneNumber))
	if err != nil {
		log.Fatal(err.Error())
	}
	if !match {
		return errors.New("phone number must be in E.164 format")
	}
	return nil
}