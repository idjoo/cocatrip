package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"jaytaylor.com/html2text"
)

type Config struct {
	Title string `yaml:"title"`
	Quote struct {
		Text   string `yaml:"text"`
		Author string `yaml:"author"`
	} `yaml:"quote"`
	Avatar    string `yaml:"avatar"`
	Name      string `yaml:"name"`
	AltName   string `yaml:"alt-name"`
	Desc      string `yaml:"desc"`
	TechStack []struct {
		Name string `yaml:"name"`
		Icon string `yaml:"icon"`
		Url  string `yaml:"url"`
	} `yaml:"techStack"`
}

func readConfig() (Config, error) {
	var config Config

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return config, err
	}

	for i := 0; i < len(config.TechStack); i++ {
		if config.TechStack[i].Url == "" {
			icon := strings.Split(config.TechStack[i].Icon, "-")[0]
			config.TechStack[i].Url = fmt.Sprintf("https://cdn.jsdelivr.net/gh/devicons/devicon/icons/%s/%s.svg", icon, config.TechStack[i].Icon)
		}
	}

	return config, nil
}

func isHtmlOutput(ua string) bool {
	if strings.Contains(ua, "curl") || strings.Contains(ua, "wget") {
		return false
	} else {
		return true
	}
}

func getIndex(c *gin.Context) {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.TechStack[0].Url)

	if isHtmlOutput(c.Request.Header["User-Agent"][0]) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"Title":     config.Title,
			"Quote":     config.Quote,
			"Avatar":    config.Avatar,
			"Name":      config.Name,
			"AltName":   config.AltName,
			"Desc":      config.Desc,
			"TechStack": config.TechStack,
		})
	} else {
		t, err := template.ParseFiles("./templates/index.tmpl.html")
		if err != nil {
			log.Fatal(err)
		}

		var buffer bytes.Buffer
		err = t.Execute(&buffer, config)
		html := buffer.String()

		output, err := html2text.FromString(html, html2text.Options{PrettyTables: true})
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, output)
	}
}

func redirect(url string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, url)
	}

	return gin.HandlerFunc(fn)
}

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	r.LoadHTMLGlob("./templates/*.tmpl.html")
	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.GET("/", getIndex)
	r.GET("/github", redirect("https://github.com/cocatrip/"))
	r.GET("/showwcase", redirect("https://cocatrip.showwcase.com/"))

	r.Run()
}
