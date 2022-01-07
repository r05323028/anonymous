package mq

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
	"github.com/r05323028/anonymous/db"
	pb "github.com/r05323028/anonymous/proto"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

type PostMessageHandler struct {
	conn *gorm.DB
}

func (h *PostMessageHandler) HandleMessage(m *nsq.Message) error {
	// process message
	p := &pb.Post{}
	err := proto.Unmarshal(m.Body, p)
	if err != nil {
		log.Fatal(err)
	}
	post := &db.Post{
		Name: p.Name,
		Author: db.Author{
			Name: p.Author,
		},
	}
	h.conn.Create(&post)
	h.conn.Commit()
	log.Println(post.ID)

	return nil
}

func RunConsumer(host string, port string, topic string, channel string, dbHost string, dbPort string, dbUser string, dbPassword string, dbName string) {
	// set db
	conn := db.NewDB(dbHost, dbPort, dbUser, dbPassword, dbName)

	// set consumer
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(&PostMessageHandler{
		conn,
	})
	dsn := fmt.Sprintf("%v:%v", host, port)
	err = consumer.ConnectToNSQLookupd(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}
