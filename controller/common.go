package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cocatrip/cocatrip/models"
	"gopkg.in/yaml.v3"
)

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
				config.Certs[i].Url = fmt.Sprintf("https://www.credly.com/badges/%s/public_url", config.Certs[i].Id)
			}
		}
	}

	for i := 0; i < len(config.SocialMedia); i++ {
    switch config.SocialMedia[i].Name {
    case "github":
      config.SocialMedia[i].Icon = "https://cdn-icons-png.flaticon.com/512/733/733609.png"
    case "instagram":
      config.SocialMedia[i].Icon = "https://cdn-icons-png.flaticon.com/512/3955/3955024.png"
    case "facebook":
      config.SocialMedia[i].Icon = "https://cdn-icons-png.flaticon.com/512/145/145802.png"
    case "twitter":
      config.SocialMedia[i].Icon = "https://cdn-icons-png.flaticon.com/512/3670/3670151.png"
    case "showwcase":
      config.SocialMedia[i].Icon = "https://github.com/Showwcase.png"
    case "reddit":
      config.SocialMedia[i].Icon = "https://cdn-icons-png.flaticon.com/512/3670/3670226.png"
    case "telegram":
      config.SocialMedia[i].Icon = "https://cdn-icons-png.flaticon.com/512/2111/2111646.png"
    }

    config.SocialMedia[i].Url = fmt.Sprintf("https://cocatrip.xyz/%s", config.SocialMedia[i].Name)
	}

	return config, nil
}

func isHtmlOutput(ua string) bool {
	if strings.Contains(ua, "curl") ||
		strings.Contains(ua, "Wget") ||
		strings.Contains(ua, "Go-http-client") {
		return false
	} else {
		return true
	}
}

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
