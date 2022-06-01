package helpers

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
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

// This function create XML-config for Allure report

func CfgToXML(path string) error {
	type parameter struct {
		Key   string `xml:"key"`
		Value string `xml:"value"`
	}

	type environment struct {
		Parameters []parameter `xml:"parameter"`
	}

	data := &environment{}
	data.Parameters = append(data.Parameters, parameter{Key: "APIHost", Value: "127.0.0.1"})
	data.Parameters = append(data.Parameters, parameter{Key: "APIPort", Value: "8080"})
	data.Parameters = append(data.Parameters, parameter{Key: "gRPCPort", Value: "8082"})

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	xlmWriter := io.Writer(file)

	enc := xml.NewEncoder(xlmWriter)

	enc.Indent("  ", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

// This function generates files for Allure trends and replaces them

func GenAllureHistory() {
	cmd := exec.Command("allure", "generate", "--clean", "./allure-results")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Generate report err:%v", err)
	}
	cmd = exec.Command("cp", "-R", "./allure-report/history", "./allure-results")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Copy report err:%v", err)
	}
}
