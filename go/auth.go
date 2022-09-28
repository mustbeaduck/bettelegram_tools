package main

import (
	"fmt"
	"github.com/Arman92/go-tdlib"
)


func getClient() *tdlib.Client {

	tdlib.SetLogVerbosityLevel(1)

	// Create new instance of client
	client := tdlib.NewClient(tdlib.Config{
		APIID:               "11280565",
		APIHash:             "5561e9f775bae9639412059f1cd4345d",
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
		UseMessageDatabase:  true,
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseTestDataCenter:   false,
		DatabaseDirectory:   "/tmp/tdlib-db",
		FileDirectory:       "/tmp/tdlib-files",
		IgnoreFileNames:     false,
	})

    return client
}

func authLoop(client *tdlib.Client) {

	for {       
		currentState, _ := client.Authorize()
		if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPhoneNumberType {
			fmt.Print("Enter phone: ")
			var number string
			fmt.Scanln(&number)
			_, err := client.SendPhoneNumber(number)
			if err != nil {
				fmt.Printf("Error sending phone number: %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitCodeType {
			fmt.Print("Enter code: ")
            var code string 
			fmt.Scanln(&code)
			_, err := client.SendAuthCode(code)         
			if err != nil {
				fmt.Printf("Error sending auth code : %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitPasswordType {
			fmt.Print("Enter Password: ")
			var password string
			fmt.Scanln(&password)
			_, err := client.SendAuthPassword(password)
			if err != nil {
				fmt.Printf("Error sending auth password: %v", err)
			}
		} else if currentState.GetAuthorizationStateEnum() == tdlib.AuthorizationStateReadyType {
            fmt.Print("authorization succes\n")
			break
		}
	}
}
