package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Create a producer
	brokerAddress := "localhost:9092"
	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, nil)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer: ", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln("Failed to close Sarama producer: ", err)
		}
	}()

	// Define the number of messages you want to send
	numMessages := 5
	for i := 1; i <= numMessages; i++ {

		// Create Kafka message
		message := &sarama.ProducerMessage{
			Topic: "sample-topic",
			Value: sarama.StringEncoder(fmt.Sprintf("Message number %d", i)),
		}

		// Send the message to Kafka
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to send message: %s\n", err)
		} else {
			log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		}

	}
}
