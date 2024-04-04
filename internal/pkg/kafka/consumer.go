package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

type Consumer struct {
	consumer sarama.Consumer
}

type Closer func() error

func New(conn string) (*Consumer, Closer) {
	config := sarama.Config{}
	config.Version = sarama.V2_7_0_0
	consumer, err := sarama.NewConsumer([]string{conn}, nil)
	if err != nil {
		log.Fatalf("failed while starting consumer: %s", err.Error())
	}

	closer := func() error {
		return consumer.Close()
	}

	return &Consumer{
		consumer: consumer,
	}, closer
}

func (c *Consumer) Consume(topic string, messages chan<- Message) error {
	partConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	for {
		select {
		case payload := <-partConsumer.Messages():
			var myMsg Message
			if err := json.Unmarshal(payload.Value, myMsg); err != nil {
				return err
			}
			messages <- myMsg
		}
	}
}
