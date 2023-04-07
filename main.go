package main

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"message-consumer/pkg/kafka"
)

func main() {
	kafkaProcessor()
}

func kafkaProcessor() {
	msgChanKafka := make(chan *ckafka.Message)
	topics := []string{"test"}
	server := "host.docker.internal:9094"
	fmt.Println("Kafka consumer has been started")
	go kafka.Consume(topics, server, msgChanKafka)
	for msg := range msgChanKafka {
		fmt.Println(string(msg.Value))
	}
}
