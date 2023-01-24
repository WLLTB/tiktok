package constant

const VideoCount = 30

const (
	RedisAddress  = "localhost:6379"
	RedisPassword = ""
	RedisDb       = 0
)

const (
	MysqlUsername = "root"
	MysqlPassword = "root"
	MysqlHost     = "127.0.0.1"
	MysqlPort     = "3306"
	MysqlDbname   = "tiktok"
)

const (
	OssEndpoint        = "ossEndpoint"
	OssAccessKeyId     = "ossAccessKeyId"
	OssAccessKeySecret = "ossAccessKeySecret"
	OssBucketUrl       = "ossBucketUrl"
)

const (
	RabbitMQURL  = "amqp://guest:guest@localhost:5672/"
	ExchangeName = "amq.topic"
	ExchangeType = "topic"
)

const PORT = ":9999"

const CoverSuffix = "?x-oss-process=video/snapshot,t_10000,m_fast"

const VideoFormat = ".mp4"

const (
	DemoQueue = "DEMO-QUEUE"
	DemoTopic = "DEMO-TOPIC"
)
