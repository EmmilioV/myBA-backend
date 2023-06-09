package messaging

import (
	"log"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type Connection struct {
	RabbitmqConnection *rabbitmq.Conn
}

func NewConnection(
	settings *Settings,
) *Connection {
	conn, err := rabbitmq.NewConn(
		settings.Url,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &Connection{
		RabbitmqConnection: conn,
	}
}
