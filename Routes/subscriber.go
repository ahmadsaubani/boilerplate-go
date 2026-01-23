package Routes

import (
	"boilerplate-go/Controller"
	"boilerplate-go/DTO"
	"boilerplate-go/Services/RabbitMQ"
	subcriberConfig "boilerplate-go/Services/RabbitMQ/subscriberConfig"
)

type subscriber struct {
	Config   subcriberConfig.SubscriberConfig
	Service  RabbitMQ.SubscriberService
	Consumer Controller.ControllerConsumer
}

type SubscriberRoutes interface {
	CollectSubscriber()
}

// collection to setup your subscription

func (s subscriber) CollectSubscriber() {
	s.Service.Add(s.Config.ExampleConfig.ConfigsSubsExample, s.Consumer.ExampleConsumer.ExampleSubscriberController)
	s.Service.ExecuteService()
}

// constructor for consumer
func NewConsumer(Service RabbitMQ.SubscriberService, u DTO.Utilities) SubscriberRoutes {
	return subscriber{
		Service: Service,
		Config: subcriberConfig.SubscriberConfig{
			ExampleConfig: subcriberConfig.SubsConfig{},
		},
		Consumer: Controller.NewConsumerController(u),
	}
}
