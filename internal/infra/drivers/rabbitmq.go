package drivers

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitQMDriver struct {
	Connection *amqp.Connection
}

func NewRabbitQMDriver(username string, password string, ipAddress string, port string) *RabbitQMDriver {

	fmt.Printf("Connecting to RabbitMQ in address %s:%s \n", ipAddress, port)

	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + ipAddress + ":" + port)
	FailOnError(err, "Failed to connect to RabbitMQ")

	fmt.Printf("Connected to RabbitMQ in %s:%s with success!\n", ipAddress, port)

	return &RabbitQMDriver{
		Connection: conn,
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (d *RabbitQMDriver) Close() {
	d.Connection.Close()
}
