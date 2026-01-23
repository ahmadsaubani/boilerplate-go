package RabbitMQ

import (
	engine_config "boilerplate-go/Config"
	subcriberConfig "boilerplate-go/Services/RabbitMQ/subscriberConfig"

	"github.com/streadway/amqp"
)

type ConsumerLogic func(data amqp.Delivery, rabbitChannel *amqp.Channel)

type RabbitConfig struct {
	Config               *engine_config.RabbitConfig
	connectionSubscriber *amqp.Connection
	connectionRepublish  *amqp.Connection
	Subscriber           rabbit
}

type RabbitServices interface {
	InitConnectionSubscriber(ApplicationName string)
	InitConnectionPublisher(ApplicationName string)
	GetServices() ServiceInterface
	GetServicesSubscriber() ServiceInterface
	Consumer() SubscriberService
}

type ServiceInterface interface {
	Publish(key RabbitKey, exhange RabbitExhangeName, data interface{}) error
	PublishRpc(key RabbitKey, exhange RabbitExhangeName, data interface{}) (msg []byte, err error)
	ReplyRpc(delivery amqp.Delivery, rabbitChannel *amqp.Channel, data interface{}) (err error)
}

func (r RabbitConfig) Consumer() SubscriberService {
	return &rabbit{connection: r.connectionSubscriber}
}

func (r RabbitConfig) GetServicesSubscriber() ServiceInterface {
	return &service{Con: r.connectionSubscriber}
}

func (r RabbitConfig) GetServices() ServiceInterface {
	return &service{Con: r.connectionRepublish}
}

type service struct {
	Con *amqp.Connection
}

type rabbit struct {
	connection           *amqp.Connection
	Config               *subcriberConfig.SubsConfig
	ControllerSubscriber func(data amqp.Delivery, rabbitChannel *amqp.Channel)
	Channel              *amqp.Channel
	Subscriber           *rabbit
}
