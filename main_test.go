package main

import (
	"testing"

	"github.com/blackieops/synonym/config"
)

func TestBuildTarget(t *testing.T) {
	conf := &config.Config{TargetBaseURL: "git.example.com/~myuser"}
	res := buildTarget(conf, "dotfiles")
	if res != "https://git.example.com/~myuser/dotfiles" {
		t.Errorf("buildTarget generated unexpected URL: %v", res)
	}
}

func TestBuildSource(t *testing.T) {
	conf := &config.Config{Hostname: "go.example.net"}
	res := buildSource(conf, "dotfiles")
	if res != "go.example.net/dotfiles" {
		t.Errorf("buildSource generated unexpected URL: %v", res)
	}
}
