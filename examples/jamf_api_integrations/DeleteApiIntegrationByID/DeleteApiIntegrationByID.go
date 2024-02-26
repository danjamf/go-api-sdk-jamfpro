package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	integrationID := 1 // Replace with the actual API Integration ID

	if err := client.DeleteApiIntegrationByID(integrationID); err != nil {
		fmt.Println("Error deleting API Integration:", err)
		return
	}

	fmt.Printf("API Integration with ID %d deleted successfully\n", integrationID)
}
