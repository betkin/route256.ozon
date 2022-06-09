package helpers

import (
	"log"
	"net/http"
	"net/url"
)

// IsAlive checks the health status of the tests object
func IsAlive(checkurl url.URL) {
	response, err := http.Get(checkurl.String())
	if err != nil {
		log.Fatalf("Service is not ready: %v", err.Error())
	}
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %v", response.StatusCode)
	}
}
