package config

import (
	"os"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// The hostname of this service. eg., "go.b8s.dev"
	Hostname string `yaml:"hostname"`

	// The FQDN to the Git source base. eg., "https://github.com/blackieops"
	TargetBaseURL string `yaml:"target_base_url"`

	// The name of the default Git branch. eg., "main"
	DefaultBranchName string `yaml:"default_branch_name"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, openErr := os.Open(configPath)
	defer file.Close()

	if openErr == nil {
		d := yaml.NewDecoder(file)
		if err := d.Decode(&config); err != nil {
			return nil, err
		}
	}

	if value, exists := os.LookupEnv("HOSTNAME"); exists {
		config.Hostname = value
	}

	if value, exists := os.LookupEnv("TARGET_BASE_URL"); exists {
		config.TargetBaseURL = value
	}

	if value, exists := os.LookupEnv("DEFAULT_BRANCH_NAME"); exists {
		config.DefaultBranchName = value
	}

	if config.DefaultBranchName == "" || config.TargetBaseURL == "" || config.Hostname == "" {
		return nil, fmt.Errorf("Required configuration was not met. Please ensure you have provided a config file or all environment variables.")
	}

	return config, nil
}
