package main

import (
	"github.com/cocatrip/cocatrip/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// templating options n shit
	router.Delims("{{", "}}")
	router.LoadHTMLGlob("./templates/*.tmpl.html")

	// favicon
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	// static folder
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")
	router.Static("/sh", "./static/sh")

	// routing
	router.GET("/", controller.GetIndex)
	router.GET("/config", controller.GetConfig)

	// redirect url
	router.GET("/github", controller.Redirect("https://github.com/cocatrip/"))
	router.GET("/reddit", controller.Redirect("https://reddit.com/u/cocatrip/"))
	router.GET("/instagram", controller.Redirect("https://instagram.com/adr_vian/"))
	router.GET("/showwcase", controller.Redirect("https://cocatrip.showwcase.com/"))

  // serve
	router.Run(":8080")
}
