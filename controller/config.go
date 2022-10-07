package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) {
	config, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

  c.JSON(http.StatusOK, config)
}
