package broker

import (
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/viper"
)

const GROUP_ID = "myGroup"

type Kafka struct {
	Consumer *kafka.Consumer
	Producer *kafka.Producer
}

func NewKafka() (*Kafka, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.host"),
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while create producer: %s/n", err.Error()))
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("kafka.host"),
		"group.id":          GROUP_ID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while create consumer: %s/n", err.Error()))
	}

	return &Kafka{
		Producer: producer,
		Consumer: consumer,
	}, nil
}
