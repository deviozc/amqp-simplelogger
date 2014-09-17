package config

const (
	// full url to rabbitmq server
	// eg. amqp://logger:logger@localhost:5672/
	RABBITMQ_URL   = "amqp://username:password@localhost:5672/"
	RABBITMQ_QUEUE = "logger-default"
)
