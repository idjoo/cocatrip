package main

import (
	"github.com/cocatrip/cocatrip/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	r.LoadHTMLGlob("./templates/*.tmpl.html")
	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/sh", "./static/sh")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.GET("/", controller.GetIndex)
	r.GET("/github", controller.Redirect("https://github.com/cocatrip/"))
	r.GET("/showwcase", controller.Redirect("https://cocatrip.showwcase.com/"))

	r.Run()
}
