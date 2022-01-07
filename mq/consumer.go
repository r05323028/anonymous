package mq

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

type MessageHandler struct{}

func (h *MessageHandler) HandleMessage(m *nsq.Message) error {
	// process message
	log.Println(string(m.Body))

	return nil
}

func RunConsumer(host string, port string, topic string, channel string) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)

	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(&MessageHandler{})
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
