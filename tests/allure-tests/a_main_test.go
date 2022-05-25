package atests

import (
	"fmt"
	"github.com/ozonmp/act-device-api/tests/allure-tests/config"
	"github.com/ozonmp/act-device-api/tests/allure-tests/internal/helpers"
	"log"
	"net/url"
	"testing"
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

	err = helpers.CfgToXML("./allure-results/environment.xml")
	if err != nil {
		log.Fatalf("Environment err:%v", err)
	}

	helpers.GenAllureHistory()
}
