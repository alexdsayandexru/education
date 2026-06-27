package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consume(host string, topic string, group string) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("%s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func main() {
	consume("localhost:9092", "test1", "group1")
}
