package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"twilio_bulk_sms/helper"
	"twilio_bulk_sms/model"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/messages", smsHandler)

	port := ":8000"

	fmt.Printf("Starting server at port%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
  err := godotenv.Load(".env.local")
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func smsHandler(writer http.ResponseWriter, request *http.Request) {

	var bulkSMSRequest model.SMSRequest

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&bulkSMSRequest)

	err := bulkSMSRequest.Validate()
	if err != nil {
		returnResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}
	
	res, err := helper.BulkSMS(bulkSMSRequest)
	if err != nil {
		returnResponse(writer, err.Error(), http.StatusBadRequest)
		return
	}

	returnResponse(writer, res, http.StatusOK)
}


func returnResponse(writer http.ResponseWriter, message string, httpStatusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	writer.Write(jsonResp)
}