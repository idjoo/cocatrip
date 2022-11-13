package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

  c.HTML(http.StatusOK, "index.html", gin.H{
    "Config": config,
  })
}
