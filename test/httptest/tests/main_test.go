package test

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	test_config "gitlab.ozon.dev/betkin/device-api/test/httptest/internal/config"
	"gitlab.ozon.dev/betkin/device-api/test/httptest/internal/helpers"
)

func TestMain(m *testing.M) {
	fmt.Println("This is HTTP test suite for device-api")
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
