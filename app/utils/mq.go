package utils

import (
	"github.com/streadway/amqp"
	"log"
	"tiktok/app/constant"
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
