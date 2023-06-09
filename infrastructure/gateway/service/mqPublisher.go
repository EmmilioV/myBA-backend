package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/wagslane/go-rabbitmq"
	"go.mod/domain/common"
	"go.mod/domain/service/gateway"
	"go.mod/infrastructure/messaging"
)

type MQPublisher struct {
	rabbitMQPublisher *rabbitmq.Publisher
}

func NewMQPublisher(
	mqConnection *messaging.Connection,
) gateway.IMQPublisher {
	publisher, err := rabbitmq.NewPublisher(
		mqConnection.RabbitmqConnection,
		rabbitmq.WithPublisherOptionsExchangeName("events"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &MQPublisher{
		rabbitMQPublisher: publisher,
	}
}

func (publisher *MQPublisher) ServiceUpdated(ctx context.Context, event *common.Event) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	data, err := json.Marshal(event)
	if err != nil {
		return errors.New("error converting data event" + err.Error())
	}

	return publisher.rabbitMQPublisher.PublishWithContext(
		ctxTimeout,
		data,
		[]string{"myBA-backend"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("events"),
	)
}
