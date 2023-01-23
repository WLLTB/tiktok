package constant

const (
	MysqlConnectSuccess = "Connect to MySQL success."
	MysqlConnectFailed  = "Connect to MySQL failed."
	MysqlQuerySuccess   = "MySQL query success."
	MysqlQueryFailed    = "MySQL query failed."
	MysqlExecSuccess    = "MySQL exec success."
	MysqlExecFailed     = "MySQL exec failed."
)

const (
	RedisConnectSuccess = "Connect to Redis success."
	RedisConnectFailed  = "Connect to Redis failed."
	RedisSetSuccess     = "Redis set success."
	RedisSetFailed      = "Redis set failed."
	RedisGetSuccess     = "Redis get success."
	RedisGetFailed      = "Redis get failed."
	RedisDelSuccess     = "Redis delete success."
	RedisDelFailed      = "Redis delete failed."
)

const (
	RabbitmqConnectSuccess       = "Connect to RabbitMQ success."
	RabbitmqConnectFailed        = "Connect to RabbitMQ failed."
	RabbitmqChannelOpenSuccess   = "Open RabbitMQ channel success."
	RabbitmqChannelDeclareFailed = "Declare RabbitMQ channel failed."
	RabbitmqChannelOpenFailed    = "Open RabbitMQ channel failed."
	RabbitmqQueueDeclareSuccess  = "Declare RabbitMQ queue success."
	RabbitmqQueueDeclareFailed   = "Declare RabbitMQ queue failed."
	RabbitmqQueueBindSuccess     = "Bind RabbitMQ queue success."
	RabbitmqQueueBindFailed      = "Bind RabbitMQ queue failed."

	RabbitmqPublishSuccess = "RabbitMQ publish success."
	RabbitmqPublishFailed  = "RabbitMQ publish failed."

	RabbitmqConsumeSuccess = "RabbitMQ consume success."
	RabbitmqConsumeFailed  = "RabbitMQ consume failed."
)

const OssConnectError = "Oss Connect Error."

const LogPrefix = "[GIN_LOG] "
