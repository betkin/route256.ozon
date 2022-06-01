package helpers

import (
	"log"
	"net/http"
	"net/url"
)

// This function checks the health status of the test object

func IsAlive(checkurl url.URL) {
	response, err := http.Get(checkurl.String())
	if err != nil {
		log.Fatalf("Service is not ready: %v", err.Error())
	}
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %v", response.StatusCode)
	}
}
