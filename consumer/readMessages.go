package main

import (
	"github.com/IBM/sarama"
	"log"
)

func main() {

	// Create a Consumer
	brokerAddress := "localhost:9092"
	consumer, err := sarama.NewConsumer([]string{brokerAddress}, nil)
	if err != nil {
		log.Fatalln("Failed to create Sarama consumer:", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama consumer:", err)
		}
	}()

	// Specify the topic to consume from
	topic := "sample-topic"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest) //'OffsetNewest' for only new message
	if err != nil {
		log.Fatalln("Failed to create partition consumer:", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln("Failed to close partition consumer:", err)
		}
	}()

	// Start consuming messages
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Received message: %s\n", msg.Value)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %s\n", err.Err)
		}
	}

}
