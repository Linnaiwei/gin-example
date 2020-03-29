package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//PingGet A handler for /ping
func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})
}