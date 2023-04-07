package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, server string, msgCan chan *ckafka.Message) {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "goapp",
		"auto.offset.reset": "earliest",
	}

	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgCan <- msg
		}
	}
}
