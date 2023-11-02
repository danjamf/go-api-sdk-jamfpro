package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
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
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define a policy category for testing
	policyCategory := "YourPolicyCategory"

	// Call GetPolicyByCategory function
	policies, err := client.GetPolicyByCategory(policyCategory)
	if err != nil {
		log.Fatalf("Error fetching policies by category: %v", err)
	}

	// Pretty print the policies in XML
	policiesXML, err := xml.MarshalIndent(policies, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling policies data: %v", err)
	}
	fmt.Println("Fetched Policies by Category:\n", string(policiesXML))
}
