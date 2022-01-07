package cmd

import (
	"fmt"
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
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		topic, _ := cmd.Flags().GetString("topic")
		channel, _ := cmd.Flags().GetString("channel")
		dbHost, _ := cmd.Flags().GetString("dbHost")
		dbPort, _ := cmd.Flags().GetString("dbPort")
		dbUser, _ := cmd.Flags().GetString("dbUser")
		dbPassword, _ := cmd.Flags().GetString("dbPassword")
		dbName, _ := cmd.Flags().GetString("dbName")
		mq.RunConsumer(host, port, topic, channel, dbHost, dbPort, dbUser, dbPassword, dbName)
	},
}

var pushMessage = &cobra.Command{
	Use:   "push",
	Short: "Push message",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		topic, _ := cmd.Flags().GetString("topic")
		message, _ := cmd.Flags().GetString("message")
		config := nsq.NewConfig()
		dsn := fmt.Sprintf("%v:%v", host, port)
		producer, err := nsq.NewProducer(dsn, config)
		if err != nil {
			log.Fatal(err)
		}
		messageBody := []byte(message)

		// Synchronously publish a single message to the specified topic.
		// Messages can also be sent asynchronously and/or in batches.
		err = producer.Publish(topic, messageBody)
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

	// consumer flags
	startConsumer.Flags().String("host", "nsqlookupd", "NSQ Daemon host")
	startConsumer.Flags().String("port", "4161", "NSQ Daemon port")
	startConsumer.Flags().String("topic", "anonymous", "NSQ Topic")
	startConsumer.Flags().String("channel", "general", "NSQ Channel")
	startConsumer.Flags().String("dbUser", "anonymous", "Database User")
	startConsumer.Flags().String("dbPassword", "example", "Database Password")
	startConsumer.Flags().String("dbHost", "mysql", "Database Host")
	startConsumer.Flags().String("dbPort", "3306", "Database Port")
	startConsumer.Flags().String("dbName", "anonymous", "Database Name")

	// producer flags
	pushMessage.Flags().String("host", "nsqd", "NSQ Daemon host")
	pushMessage.Flags().String("port", "4150", "NSQ Daemon port")
	pushMessage.Flags().String("topic", "anonymous", "Topic")
	pushMessage.Flags().String("message", "", "Message")
}
