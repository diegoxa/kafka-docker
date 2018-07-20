package main

import (
	"fmt"

	"bufio"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

var bootstrapServers = "10.122.0.108:29092"

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "topic"
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(scanner.Text()),
		}, nil)
	}

	//// Produce messages to topic (asynchronously)
	//topic := "diego.data_mart.url_content.crawl"
	//for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
	//	p.Produce(&kafka.Message{
	//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	//		Value:          []byte(word),
	//	}, nil)
	//}

	// Wait for message deliveries
	p.Flush(15 * 1000)
}
