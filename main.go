package main

import (
	"embed"
	"flag"
	"html/template"
	"net/http"

	"github.com/blackieops/synonym/config"
	"github.com/gin-gonic/gin"
)

//go:embed tmpl/*
var tmplFS embed.FS

var configPath = flag.String("config", "config.yaml", "Path to configuration file.")

func main() {
	flag.Parse()

	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	tmpls := template.Must(template.New("").ParseFS(tmplFS, "tmpl/*"))
	r.SetHTMLTemplate(tmpls)
	r.GET("/*importPath", handleGetRepo(conf))
	r.Run(":6969")
}

func handleHealthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func handleGetRepo(conf *config.Config) func(c *gin.Context) {
	return func(c *gin.Context) {
		name := c.Param("importPath")[1:]

		if name == "_healthz" {
			handleHealthz(c)
			return
		}

		target := buildTarget(conf, name)
		if c.Query("go-get") == "1" {
			c.HTML(http.StatusOK, "go-get.html", gin.H{
				"Source":            buildSource(conf, name),
				"Target":            target,
				"DefaultBranchName": conf.DefaultBranchName,
			})
			return
		}
		c.Redirect(http.StatusPermanentRedirect, target)
	}
}

func buildTarget(config *config.Config, repo string) string {
	return "https://" + config.TargetBaseURL + "/" + repo
}

func buildSource(config *config.Config, repo string) string {
	return config.Hostname + "/" + repo
}
