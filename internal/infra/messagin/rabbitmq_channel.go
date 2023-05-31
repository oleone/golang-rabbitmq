package messagin

import (
	"fmt"

	"github.com/oleone/golang-rabbitmq/internal/infra/drivers"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqChannel struct {
	Channel *amqp.Channel
}

func NewRabbitMqChannel(driver *drivers.RabbitQMDriver) *RabbitMqChannel {
	fmt.Println("Opening the channel with RabbitMQ")

	ch, err := driver.Connection.Channel()
	drivers.FailOnError(err, "Failed to open a channel")

	fmt.Println("Channel by RabbitMQ is opened!")

	return &RabbitMqChannel{
		Channel: ch,
	}
}

func (c *RabbitMqChannel) Close() {
	c.Channel.Close()
}
