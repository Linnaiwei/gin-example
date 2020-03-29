package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-example-with-gin.com/httpd/handler"
	"github.com/go-example-with-gin.com/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.Run() // listen and serve on 0.0.0.0:8080
}