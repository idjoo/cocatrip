package main

import (
	"github.com/cocatrip/cocatrip/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	// declare router
	router := gin.Default()

	// templating options n shit
	router.Delims("{{", "}}")
	router.LoadHTMLGlob("./templates/*.html")

	// favicon
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	// static folder
	router.Static("/css", "./static/css")

	// routing
	router.GET("/", controller.GetIndex)
	router.GET("/config", controller.GetConfig)

	// redirect url
	router.GET("/github", controller.Redirect("https://github.com/cocatrip"))
	router.GET("/twitter", controller.Redirect("https://twitter.com/cocatrip"))
	router.GET("/reddit", controller.Redirect("https://reddit.com/u/cocatrip"))
	router.GET("/instagram", controller.Redirect("https://instagram.com/adr_vian"))
	router.GET("/showwcase", controller.Redirect("https://cocatrip.showwcase.com"))
	router.GET("/facebook", controller.Redirect("https://fb.com/adrianus.vian"))
	router.GET("/telegram", controller.Redirect("https://t.me/adr_vian"))
	router.GET("/dotfiles", controller.Redirect("https://github.com/cocatrip/dotfiles"))

	// serve
	router.Run(":8080")
}
