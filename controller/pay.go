package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPay(c *gin.Context) {
	c.HTML(http.StatusOK, "pay.html", gin.H{})
}
