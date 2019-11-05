package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var bootstrapServers = "10.120.2.96:29092"

func main() {
	// library librdkafka is needed
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          "group1",
		"auto.offset.reset": "earliest",
		//"auto.offset.reset": "largest",
		//"auto.offset.reset": "smallest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"topic"}, nil)
	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}
