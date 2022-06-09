package test_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config describes the configuration of the tests object
type Config struct {
	APIHost       string `envconfig:"API_HOST" default:"127.0.0.1"`
	APIPort       string `envconfig:"API_PORT" default:"8080"`
	GRPCPort      string `envconfig:"GRPC_PORT" default:"8082"`
	LivecheckPort string `envconfig:"LIVECHECK_PORT" default:"8000"`
	LivecheckURI  string `envconfig:"LIVECHECK_URI" default:"live"`
}

// GetConfig create configuration object
func GetConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	return config, err
}

// GetAPIURL returns URL for connection to API
func GetAPIURL(cfg Config) string {
	return fmt.Sprintf("http://%s:%s", cfg.APIHost, cfg.APIPort)
}

// GetGrpcURL returns URL to gRPC connection to API
func GetGrpcURL(cfg Config) string {
	return fmt.Sprintf("%s:%s", cfg.APIHost, cfg.GRPCPort)
}
