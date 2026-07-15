package producer

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func (p *Producer) Close() error {
	return p.producer.Close()
}

func NewProducer(addr string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		log.Fatalf("error creating producer: %v", err)
	}
	return &Producer{producer: producer}, nil

}

func (p *Producer) SendMessage(topic, content string) error {
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(content),
	}

	partition, offset, err := p.producer.SendMessage(message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}
