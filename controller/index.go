package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/cocatrip/cocatrip/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"jaytaylor.com/html2text"
)

func getCredlyData(id string) (models.Credly, error) {
	var credly models.Credly

	client := &http.Client{}

  url := fmt.Sprintf("https://www.credly.com/api/v1/public_badges/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return credly, err
	}

	res, err := client.Do(req)
	if err != nil {
		return credly, err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return credly, err
	}

	err = json.Unmarshal(bodyBytes, &credly)
	if err != nil {
		return credly, err
	}

  return credly, err
}

func readConfig() (models.Config, error) {
	var config models.Config

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

	for i := 0; i < len(config.Certs); i++ {
		if config.Certs[i].Id != "" {
			switch config.Certs[i].Provider {
			case "credly":
        credly, err := getCredlyData(config.Certs[i].Id)
        if err != nil {
          return config, err
        }
        
        config.Certs[i].Name = credly.Data.BadgeTemplate.Name
        config.Certs[i].Image = credly.Data.BadgeTemplate.ImageURL
        config.Certs[i].Issuer = credly.Data.Issuer.Summary
			}
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

func GetIndex(c *gin.Context) {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	if isHtmlOutput(c.Request.Header["User-Agent"][0]) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"Title":     config.Title,
			"Quote":     config.Quote,
			"Avatar":    config.Avatar,
			"Name":      config.Name,
			"Alias":     config.Alias,
			"Desc":      config.Desc,
			"TechStack": config.TechStack,
			"Certs":     config.Certs,
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
