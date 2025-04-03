package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/plaid/plaid-go/plaid"
)

func GetPlaidLinkToken() *plaid.APIClient {

    // retrieve credentials from environment variables
    clientID := os.Getenv("PLAID_CLIENT_ID")
    secret := os.Getenv("PLAID_SECRET")
    environment := os.Getenv("PLAID_ENVIRONMENT")

    if clientID == "" || secret == "" || environment == ""{
        log.Fatal("PLAID_CLIENT_ID, PLAID_SECRET, or PLAID_ENVIRONMENT environment variable is missing.")
    }

    // create a new config for the Plaid API client
    configuration := plaid.NewConfiguration()

    // Set headers with the Plaid Client ID and Secret
    configuration.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
    configuration.AddDefaultHeader("PLAID-SECRET", secret)

    // use the correct environment
    switch environment {
    case "sandbox":
        configuration.UseEnvironment(plaid.Sandbox)
    case "development":
        configuration.UseEnvironment(plaid.Development)
    case "production":
        configuration.UseEnvironment(plaid.Production)
    default:
        log.Fatal("Invalid PLAID_ENVIRONMENT value. use sandbox, development, or production.")
    }
    // create and return the API client
    return plaid.NewAPIClient(configuration)
}

func main() {

    client := GetPlaidLinkToken()
    context := context.Background()
    // ex API call for testing 
    response := client.PlaidApi.InstitutionsGet(context)

    fmt.Printf("Response: %v\n", response)
}
