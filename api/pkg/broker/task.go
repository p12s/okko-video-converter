package broker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/p12s/okko-video-converter/api/common"
	"github.com/p12s/okko-video-converter/api/pkg/repository"
	"github.com/p12s/okko-video-converter/api/utils/processing"
	"github.com/spf13/viper"
)

func Create(producer *confluentKafka.Producer, data common.VideoConvertData) error {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(common.Event{
		Type:  common.EVENT_VIDEO_CONVERT,
		Value: data,
	}); err != nil {
		return errors.New("error while task data trying encoding")
	}

	topic := viper.GetString("kafka.topic")
	if err := producer.Produce(&confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{
			Topic:     &topic,
			Partition: confluentKafka.PartitionAny,
		},
		Value: b.Bytes(),
	}, nil); err != nil {
		return errors.New(fmt.Sprintf(
			"error while send produce message to kafka: %s\n", err.Error()))
	}

	return nil
}

func Subscribe(consumer *confluentKafka.Consumer, repos *repository.Repository) {
	err := consumer.SubscribeTopics([]string{viper.GetString("kafka.topic")}, nil)
	if err != nil {
		fmt.Println("Subscribe kafka ERROR:", err.Error())
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var event common.Event
			err := json.Unmarshal(msg.Value, &event)
			if err != nil {
				fmt.Println("Unmarshal error while decode kafka event:", err.Error())
			}
			go processing.ProcessEvent(event, repos)
			fmt.Printf("Process event from kafka: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Process event from ERROR: %v (%v)\n", err, msg) // TODO логировать
		}
	}
}
