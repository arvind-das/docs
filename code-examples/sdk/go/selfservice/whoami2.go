package main

import (
	"context"
	"fmt"
	"os"

	ory "github.com/ory/client-go"
)

func main() {
	proxyPort := os.Getenv("PROXY_PORT")
	if proxyPort == "" {
		proxyPort = "4000"
	}
	configuration := ory.NewConfiguration()
	configuration.Servers = ory.ServerConfigurations{{URL: fmt.Sprintf("http://localhost:%s/.ory", proxyPort)}}
	apiClient := ory.NewAPIClient(configuration)
	cookie := "ory_session_playground=<your-session-cookie-here>"
	resp, r, err := apiClient.FrontendApi.ToSession(context.Background()).Cookie(cookie).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrontendApi.ToSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToSession`: Session
	fmt.Fprintf(os.Stdout, "Traits  %v\n", resp.Identity.Traits)
}
