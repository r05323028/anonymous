package mq

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

type MessageHandler struct{}

func (h *MessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}
	log.Println(string(m.Body))
	err := errors.New("Test")
	return err
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
