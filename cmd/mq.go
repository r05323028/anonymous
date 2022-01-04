package cmd

import (
	"log"

	"github.com/nsqio/go-nsq"
	"github.com/r05323028/anonymous/mq"
	"github.com/spf13/cobra"
)

var mqCmd = &cobra.Command{
	Use:   "mq",
	Short: "Message Queue Command Line Tools",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startConsumer = &cobra.Command{
	Use:   "start-consumer",
	Short: "Start consumer",
	Run: func(cmd *cobra.Command, args []string) {
		mq.RunConsumer(host, port)
	},
}

var pushMessage = &cobra.Command{
	Use:   "push",
	Short: "Push message",
	Run: func(cmd *cobra.Command, args []string) {
		config := nsq.NewConfig()
		producer, err := nsq.NewProducer("nsqd:4150", config)
		if err != nil {
			log.Fatal(err)
		}

		messageBody := []byte("hello")
		topicName := "anonymous"

		// Synchronously publish a single message to the specified topic.
		// Messages can also be sent asynchronously and/or in batches.
		err = producer.Publish(topicName, messageBody)
		if err != nil {
			log.Fatal(err)
		}

		// Gracefully stop the producer when appropriate (e.g. before shutting down the service)
		producer.Stop()
	},
}

func init() {
	mqCmd.AddCommand(startConsumer)
	mqCmd.AddCommand(pushMessage)
	rootCmd.AddCommand(mqCmd)
}
