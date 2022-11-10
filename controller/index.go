package controller

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	if isHtmlOutput(c.Request.Header["User-Agent"][0]) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":     config.Title,
			"Quote":     config.Quote,
			"Logo":      config.Logo,
			"Name":      config.Name,
			"Alias":     config.Alias,
			"Desc":      config.Desc,
			"TechStack": config.TechStack,
			"Certs":     config.Certs,
			"SocialMedia":     config.SocialMedia,
		})
	} else {
		t, err := template.ParseFiles("./templates/index.md")
		if err != nil {
			log.Fatal(err)
		}

		var buffer bytes.Buffer
		err = t.Execute(&buffer, config)
		output := buffer.String()

		c.String(http.StatusOK, output)
	}
}
