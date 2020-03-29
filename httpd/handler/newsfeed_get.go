package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-example-with-gin.com/platform/newsfeed"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
