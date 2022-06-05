package tests

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	test_config "gitlab.ozon.dev/betkin/device-api/tests/config"
	"gitlab.ozon.dev/betkin/device-api/tests/internal/helpers"
)

func TestMain(m *testing.M) {
	fmt.Println("This is gRPC test suite for device-api")
	cfg, err := test_config.GetConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err.Error())
	}
	checkurl := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.APIHost, cfg.LivecheckPort),
		Path:   fmt.Sprintf("/%s", cfg.LivecheckURI),
	}
	helpers.IsAlive(checkurl)
	m.Run()
}
