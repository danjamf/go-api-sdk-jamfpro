package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	concurrentRequests           = 10 // Number of simultaneous requests.
	maxConcurrentRequestsAllowed = 5  // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:          authConfig.InstanceName,
		DebugMode:             true,
		Logger:                jamfpro.NewDefaultLogger(),
		MaxConcurrentRequests: maxConcurrentRequestsAllowed,
		TokenLifespan:         defaultTokenLifespan,
		BufferPeriod:          defaultBufferPeriod,
		ClientID:              authConfig.ClientID,
		ClientSecret:          authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client := jamfpro.NewClient(config)

	// Assemble the request body for updating an account
	accountToUpdate := &jamfpro.ResponseAccount{
		Name:                "John Smith",
		DirectoryUser:       false,
		FullName:            "John Smith Updated",
		Email:               "john.smith.updated@company.com",
		EmailAddress:        "john.smith.updated@company.com",
		Enabled:             "Enabled",
		ForcePasswordChange: true,
		AccessLevel:         "Full Access",
		PrivilegeSet:        "Administrator",
		Password:            "sampleUpdated",
		Site: jamfpro.AccountDataSubsetSite{
			ID:   -1,
			Name: "None",
		},
		Privileges: jamfpro.AccountDataSubsetPrivileges{
			JSSObjects:    []string{"updatedString"},
			JSSSettings:   []string{"updatedString"},
			JSSActions:    []string{"updatedString"},
			Recon:         []string{"updatedString"},
			CasperAdmin:   []string{"updatedString"},
			CasperRemote:  []string{"updatedString"},
			CasperImaging: []string{"updatedString"},
		},
	}

	// Let's assume we are updating an account with the name "Bobby".
	accountName := "John Smith Updated2"

	// Call UpdateAccountByName function
	updatedAccount, err := client.UpdateAccountByName(accountName, accountToUpdate)

	if err != nil {
		log.Fatalf("Error updating account by name: %v", err)
	}

	// Pretty print the updated account details
	accountXML, err := xml.MarshalIndent(updatedAccount, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Updated Account Details:", string(accountXML))
}
