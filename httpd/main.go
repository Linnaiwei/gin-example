package main

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	// "github.com/go-example-with-gin.com/httpd/handler"
	"github.com/go-example-with-gin.com/platform/newsfeed"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", handler.PingGet())
	// r.Run() // listen and serve on 0.0.0.0:8080
	feed := newsfeed.New()
	fmt.Println(feed)
	feed.Add(newsfeed.Item{
		Title: "Hello",
		Post: "World",
	})
	fmt.Println(feed.GetAll())
}