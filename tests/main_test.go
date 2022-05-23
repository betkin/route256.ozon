package tests

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/ozonmp/act-device-api/tests/config"
	"github.com/ozonmp/act-device-api/tests/internal/helpers"
)

func TestMain(m *testing.M) {
	fmt.Println("This is gRPC test suite for device-api")
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err.Error())
	}
	checkurl := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.LivecheckPort),
		Path:   fmt.Sprintf("/%s", cfg.LivecheckUri),
	}
	helpers.IsAlive(checkurl)
	m.Run()
}
