package controller

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"jaytaylor.com/html2text"
)

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
