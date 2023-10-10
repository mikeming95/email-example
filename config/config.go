package config

import (
	"os"

	"github.com/jinzhu/configor"
)

type Email struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	LoginEmail    string `yaml:"loginEmail"`
	LoginPassword string `yaml:"loginPassword"`
	SenderEmail   string `yaml:"senderEmail"`
}

var EmailConfig = Email{}

func init() {
	configFilePath := "config/config.yaml"
	// check if config file exists or not exists then use default config file
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		configFilePath = "config/config.default.yaml"
	}
	configor.Load(&EmailConfig, configFilePath)
}
