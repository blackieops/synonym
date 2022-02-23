package config

import (
	"testing"
)

func TestLoadConfigValid(t *testing.T) {
	c, err := LoadConfig("internal/fixtures/valid.yaml")
	if err != nil {
		t.Errorf("Failed to load valid config: %v", err)
	}
	if c.TargetBaseURL != "github.com/blackieops" {
		t.Errorf("Parsed Port incorrectly: %v", c.TargetBaseURL)
	}
	if c.Hostname != "go.example.net" {
		t.Errorf("Parsed BaseURL incorrectly: %v", c.Hostname)
	}
	if c.DefaultBranchName != "trunk" {
		t.Errorf("Parsed BaseURL incorrectly: %v", c.DefaultBranchName)
	}
}

func TestLoadConfigInvalid(t *testing.T) {
	_, err := LoadConfig("internal/fixtures/notexist.yaml")
	if err == nil {
		t.Errorf("Somehow loaded non-existent config!")
	}
	_, err = LoadConfig("internal/fixtures/invalid.yaml")
	if err == nil {
		t.Errorf("Erroneously loaded invalid config!")
	}
	_, err = LoadConfig("internal/fixtures/useless.yaml")
	if err != nil {
		t.Errorf("Should have loaded useless but syntactically valid config!")
	}
}
