package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func produce(host string, topic string, number int) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	for i := 0; i < number; i++ {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(strconv.Itoa(i)),
		}, nil)
		fmt.Printf("%v: %v\n", topic, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	produce("localhost:9092", "test1", 200)
}
