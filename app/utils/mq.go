package utils

import (
	"github.com/streadway/amqp"
)

var RabbitMQConnection *amqp.Connection

func PublishToTopic(msg []byte, topic string) {
	ch, err := RabbitMQConnection.Channel()
	if err != nil {
		panic(err)
	}
	err = ch.ExchangeDeclare(
		"topic_exchange", // name
		"topic",          // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"topic_exchange", // exchange
		topic,            // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	if err != nil {
		panic(err)
	}
}
