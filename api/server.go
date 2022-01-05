package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/r05323028/anonymous/proto"
)

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, struct {
		Version string    `json:"version"`
		Name    string    `json:"name"`
		Ts      time.Time `json:"ts"`
	}{
		Name:    "Anonymous",
		Version: "0.1.0",
		Ts:      time.Now(),
	})
}

func main() {
	r := gin.Default()

	// post group
	postGroup := r.Group("/post")
	postGroup.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")
		author := c.PostForm("author")
		msg := &pb.Post{
			Name:   name,
			Author: author,
		}
		c.IndentedJSON(http.StatusOK, struct {
			Status string
		}{
			Status: "OK",
		})
		log.Println(msg)
	})

	// health check
	r.GET("/health", healthCheck)

	r.Run()
}
