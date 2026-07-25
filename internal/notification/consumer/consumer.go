package consumer

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}

func NewConsumer(addr string) (*Consumer, error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{addr}, config)
	if err != nil {
		return nil, fmt.Errorf("error creating producer: %v", err)
	}
	return &Consumer{consumer: consumer}, nil
}

func (c *Consumer) ConsumePartition(topic string) error {
	consumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return fmt.Errorf("failed to start partition consumer: %v", err)
	}
	for message := range consumer.Messages() {
		log.Printf("received message: %s", string(message.Value))
	}
	return nil
}
