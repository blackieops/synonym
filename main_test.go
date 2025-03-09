package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackieops/synonym/config"
	"github.com/stretchr/testify/assert"
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

func TestHandleHealthz(t *testing.T) {
	req, err := http.NewRequest("GET", "/example?go-get=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handleHealthz(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body := rr.Body.String()
	assert.Equal(t, body, "ok\n")
}

func TestHandleGetRepoGoGet(t *testing.T) {
	conf := &config.Config{
		TargetBaseURL:     "worktree.ca/blackieops",
		DefaultBranchName: "main",
	}

	handler := handleGetRepo(conf)

	req, err := http.NewRequest("GET", "/example?go-get=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	body := rr.Body.String()
	assert.Contains(t, body, `<meta name="go-import" content="/example git https://worktree.ca/blackieops/example">`)
	assert.Contains(t, body, `name="go-source"`)
	assert.Contains(t, body, `https://worktree.ca/blackieops/example/tree/main{/dir}`)
	assert.Contains(t, body, `https://worktree.ca/blackieops/example/blob/main{/dir}/{file}#L{line}`)
}

func TestHandleGetRepoRedirect(t *testing.T) {
	conf := &config.Config{
		TargetBaseURL:     "worktree.ca/blackieops",
		DefaultBranchName: "main",
	}

	handler := handleGetRepo(conf)

	req, err := http.NewRequest("GET", "/example", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMovedPermanently, rr.Code)
	assert.Equal(t, "https://worktree.ca/blackieops/example", rr.Header().Get("location"))
}
