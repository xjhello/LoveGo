package rabbitmq

import (
	"github.com/streadway/amqp"
)

const (
	MQYRL = "amqp://admin:admin@172.31.236.43:15672/"
)

type RabbitMQ struct {
	coon      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}
