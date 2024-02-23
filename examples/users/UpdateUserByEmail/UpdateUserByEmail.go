package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:            loadedConfig.ClientOptions.LogLevel,
			LogOutputFormat:     loadedConfig.ClientOptions.LogOutputFormat,
			LogConsoleSeparator: loadedConfig.ClientOptions.LogConsoleSeparator,
			HideSensitiveData:   loadedConfig.ClientOptions.HideSensitiveData,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create a sample user to be created
	updatedUser := &jamfpro.ResourceUser{
		ID:           1,
		Name:         "AHarrison",
		FullName:     "Ashley Harrison",
		Email:        "aharrison@company.com",
		EmailAddress: "aharrison@company.com",
		PhoneNumber:  "123-555-6789",
		Position:     "Teacher",
		Sites: []jamfpro.SharedResourceSite{
			{
				ID:   -1,
				Name: "None",
			},
		},
	}

	// Replace with the email of the user you want to update
	userEmail := "name@domain.com"

	// Call UpdateUserByID to update the user
	updatedUser, err = client.UpdateUserByEmail(userEmail, updatedUser)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return
	}

	// Print the details of the updated user
	fmt.Printf("Updated User Details:\nID: %d\nName: %s\nEmail: %s\n", updatedUser.ID, updatedUser.Name, updatedUser.Email)
}
