package main

import (
	"encoding/xml"
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

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

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
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call the GetLDAPServerByNameAndUserMembershipInGroupDataSubset function
	ldapServerName := "exampleLDAPServerName"
	groupName := "exampleGroupName"
	userName := "exampleUserName"
	ldapServer, err := client.GetLDAPServerByNameAndUserMembershipInGroupDataSubset(ldapServerName, groupName, userName)
	if err != nil {
		log.Fatalf("Error retrieving LDAP server user membership in group data: %v", err)
	}

	// Process and print the response
	ldapServersXML, err := xml.MarshalIndent(ldapServer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling LDAP servers data: %v", err)
	}
	fmt.Println("Fetched LDAP Servers List:", string(ldapServersXML))
}
