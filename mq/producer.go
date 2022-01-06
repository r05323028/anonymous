package mq

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"
	pb "github.com/r05323028/anonymous/proto"
	"google.golang.org/protobuf/proto"
)

type Producer struct {
	host    string
	port    string
	topic   string
	channel string
}

func NewProducer(host string, port string, topic string, channel string) *Producer {
	return &Producer{
		host:    host,
		port:    port,
		topic:   topic,
		channel: channel,
	}
}

func (p *Producer) PushPost(m *pb.Post) {
	config := nsq.NewConfig()
	dsn := fmt.Sprintf("%v:%v", p.host, p.port)
	producer, err := nsq.NewProducer(dsn, config)
	if err != nil {
		log.Fatal(err)
	}
	messageBody, err := proto.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	err = producer.Publish(p.topic, messageBody)
	if err != nil {
		log.Fatal(err)
	}
	producer.Stop()
}
