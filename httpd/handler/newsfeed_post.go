package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-example-with-gin.com/platform/newsfeed"
	"net/http"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		_ = c.Bind(&requestBody)
		feed.Add(newsfeed.Item(requestBody))
		c.Status(http.StatusNoContent)
	}
}
