package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/userauth.json"

	// Load the client authentication details from the configuration file
	authConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client authentication configuration: %v", err)
	}

	// Configuration for the client
	config := http_client.Config{
		DebugMode: true,
		Logger:    http_client.NewDefaultLogger(),
	}

	// Create a new client instance using the loaded InstanceName
	client := http_client.NewClient(authConfig.InstanceName, config, nil)

	// Set BearerTokenAuthCredentials
	client.SetBearerTokenAuthCredentials(http_client.BearerTokenAuthCredentials{
		Username: authConfig.Username,
		Password: authConfig.Password,
	})

	// Call the ObtainToken function to get a bearer token
	err = client.ObtainToken()
	if err != nil {
		fmt.Println("Error obtaining token:", err)
		return
	}

	// Print the obtained token
	fmt.Println("Successfully obtained token:", client.Token)
}
