package utils

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"tiktok/app/constant"
	"time"
)

var RabbitMQConnection *amqp.Connection

func PublishToTopic(msg []byte, topic string) {
	ch, err := RabbitMQConnection.Channel()
	if err != nil {
		log.Println(constant.RabbitmqChannelOpenFailed)
	}
	err = ch.ExchangeDeclare(
		constant.ExchangeName, // name
		"topic",               // type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		constant.ExchangeName, // exchange
		topic,                 // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	if err != nil {
		log.Println(constant.RabbitmqPublishFailed)
	}
}

func consumeTopic(ch *amqp.Channel, topic string) {
	err := ch.ExchangeDeclare(
		constant.ExchangeName, // name
		constant.ExchangeType, // type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		log.Println(constant.RabbitmqChannelDeclareFailed)
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(constant.RabbitmqQueueDeclareFailed)
	}

	err = ch.QueueBind(
		q.Name,                // queue name
		topic,                 // routing key
		constant.ExchangeName, // exchange
		false,
		nil)
	if err != nil {
		log.Println(constant.RabbitmqQueueBindFailed)
	}
	log.Println(constant.RabbitmqQueueBindSuccess)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf(" [x] %s: %s\n", d.RoutingKey, d.Body)
			time.Sleep(time.Second)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
