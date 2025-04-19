package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/blackieops/synonym/config"
)

//go:embed tmpl/*
var tmplFS embed.FS

var tmpls = template.Must(template.New("").ParseFS(tmplFS, "tmpl/*"))

var configPath = flag.String("config", "config.yaml", "Path to configuration file.")

func main() {
	flag.Parse()

	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/_healthz", handleHealthz)
	mux.HandleFunc("/", handleGetRepo(conf))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), mux); err != nil {
		panic(err)
	}
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func handleGetRepo(conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[1:]

		target := buildTarget(conf, name)

		if r.URL.Query().Get("go-get") == "1" {
			data := struct {
				Source            string
				Target            string
				DefaultBranchName string
			}{
				Source:            buildSource(conf, name),
				Target:            target,
				DefaultBranchName: conf.DefaultBranchName,
			}
			if err := tmpls.ExecuteTemplate(w, "go-get.html", data); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		http.Redirect(w, r, target, http.StatusMovedPermanently)
	}
}

func buildTarget(config *config.Config, repo string) string {
	target := config.TargetBaseURL + "/" + repo

	for _, mapping := range config.CustomMappings {
		if strings.TrimLeft(mapping.Path, "/") == repo {
			target = mapping.Target
		}
	}

	return "https://" + target
}

func buildSource(config *config.Config, repo string) string {
	return config.Hostname + "/" + repo
}
