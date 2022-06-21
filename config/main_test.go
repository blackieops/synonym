package config

import (
	"os"
	"testing"
)

func TestLoadConfigValid(t *testing.T) {
	c, err := LoadConfig("internal/fixtures/valid.yaml")
	if err != nil {
		t.Errorf("Failed to load valid config: %v", err)
	}
	if c.TargetBaseURL != "github.com/blackieops" {
		t.Errorf("Parsed TargetBaseURL incorrectly: %v", c.TargetBaseURL)
	}
	if c.Hostname != "go.example.net" {
		t.Errorf("Parsed Hostname incorrectly: %v", c.Hostname)
	}
	if c.DefaultBranchName != "trunk" {
		t.Errorf("Parsed DefaultBranchName incorrectly: %v", c.DefaultBranchName)
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
}

func TestLoadConfigWithEnvironmentAndFile(t *testing.T) {
	os.Setenv("TARGET_BASE_URL", "github.com/alexblackie")
	c, err := LoadConfig("internal/fixtures/valid.yaml")
	if err != nil {
		t.Errorf("Failed to load valid config: %v", err)
	}
	if c.TargetBaseURL != "github.com/alexblackie" {
		t.Errorf("Parsed TargetBaseURL incorrectly: %v", c.TargetBaseURL)
	}
	if c.Hostname != "go.example.net" {
		t.Errorf("Parsed Hostname incorrectly: %v", c.Hostname)
	}
	if c.DefaultBranchName != "trunk" {
		t.Errorf("Parsed DefaultBranchName incorrectly: %v", c.DefaultBranchName)
	}
}

func TestLoadConfigWithEnvironmentAndNoFile(t *testing.T) {
	os.Setenv("TARGET_BASE_URL", "github.com/alexblackie")
	os.Setenv("HOSTNAME", "localhost")
	os.Setenv("DEFAULT_BRANCH_NAME", "trunc")
	c, err := LoadConfig("internal/fixtures/notexist.yaml")
	if err != nil {
		t.Errorf("Failed to load valid config: %v", err)
	}
	if c.TargetBaseURL != "github.com/alexblackie" {
		t.Errorf("Parsed TargetBaseURL incorrectly: %v", c.TargetBaseURL)
	}
	if c.Hostname != "localhost" {
		t.Errorf("Parsed Hostname incorrectly: %v", c.Hostname)
	}
	if c.DefaultBranchName != "trunc" {
		t.Errorf("Parsed DefaultBranchName incorrectly: %v", c.DefaultBranchName)
	}
}
