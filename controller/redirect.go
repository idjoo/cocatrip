package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(url string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, url)
	}

	return gin.HandlerFunc(fn)
}

