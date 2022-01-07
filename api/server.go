package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/r05323028/anonymous/mq"
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
	// read env
	err := godotenv.Load(".env.development")
	host := os.Getenv("NSQ_HOST")
	port := os.Getenv("NSQ_PORT")
	topic := os.Getenv("NSQ_TOPIC")
	channel := os.Getenv("NSQ_CHANNEL")
	if err != nil {
		log.Fatal(err)
	}

	// build server
	r := gin.Default()

	// post group
	postGroup := r.Group("/post")

	// create new post
	producer := mq.NewProducer(host, port, topic, channel)
	postGroup.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")
		author := c.PostForm("author")
		msg := &pb.Post{
			Name:   name,
			Author: author,
		}
		producer.PushPost(msg)
		c.IndentedJSON(http.StatusOK, struct {
			Status string
		}{
			Status: "OK",
		})
	})

	// health check
	r.GET("/health", healthCheck)

	r.Run()
}
