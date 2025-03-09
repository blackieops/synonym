package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"net/http"

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
	mux.HandleFunc("/[a-zA-Z0-9_./-]+", handleGetRepo(conf))
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), mux)
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func handleGetRepo(conf *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[1:]

		if name == "_healthz" {
			handleHealthz(w, r)
			return
		}

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

			tmpls.ExecuteTemplate(w, "go-get.html", data)

			return
		}
		http.Redirect(w, r, target, http.StatusMovedPermanently)
	}
}

func buildTarget(config *config.Config, repo string) string {
	return "https://" + config.TargetBaseURL + "/" + repo
}

func buildSource(config *config.Config, repo string) string {
	return config.Hostname + "/" + repo
}
