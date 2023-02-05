package utils

import (
	"app/constant"
	"github.com/streadway/amqp"
	"log"
)

var RabbitMQConnection *amqp.Connection

func PublishToTopic(msg []byte, topic string) {
	ch, err := RabbitMQConnection.Channel()
	if err != nil {
		log.Println(constant.RabbitmqChannelOpenFailed)
	}

	err = ch.Publish(
		constant.ExchangeName,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	if err != nil {
		log.Println(constant.RabbitmqPublishFailed)
	}
}
