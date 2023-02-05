package config

import (
	. "app/constant"
	. "app/utils"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Db *gorm.DB

func InitGormDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MysqlUsername, MysqlPassword, MysqlHost, MysqlPort, MysqlDbname)
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	log.Print(MysqlConnectSuccess)
}

func InitLog() {
	log.SetPrefix(LogPrefix)
}

// InitRedisClient 初始化 Redis 客户端
func InitRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisAddress,
		Password: RedisPassword,
		DB:       RedisDb,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Panic(RedisConnectFailed, err)
	}
	log.Print(RedisConnectSuccess)
}

// InitOssClient 初始化 Oss 客户端
func InitOssClient() {
	var err error
	OssClient, err = oss.New(OssEndpoint, OssAccessKeyId, OssAccessKeySecret)
	if err != nil {
		log.Panic(OssConnectError, err)
	}
}

func InitRabbitMQ() {
	var err error
	RabbitMQConnection, err = amqp.Dial(RabbitMQURL)
	if err != nil {
		log.Fatal(RabbitmqConnectFailed)
	}

	initDemoQueue()
}

func initDemoQueue() {
	ch, err := RabbitMQConnection.Channel()
	if err != nil {
		log.Println(RabbitmqChannelOpenFailed)
	}
	_, err = ch.QueueDeclare(
		DemoQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf(RabbitmqQueueDeclareFailed)
	}
	log.Println(RabbitmqQueueDeclareSuccess)

	err = ch.QueueBind(
		DemoQueue,
		DemoTopic,
		ExchangeName,
		false,
		nil)
	if err != nil {
		log.Fatalf(RabbitmqQueueBindFailed)
	}
	log.Println(RabbitmqQueueBindSuccess)

	messages, err := ch.Consume(
		DemoQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for message := range messages {
			log.Printf("Received a messages: %s", message.Body)
		}
	}()
}
