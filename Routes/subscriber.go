package Routes

import (
	"boilerplate-go/Controller/Subscribers"
	"boilerplate-go/DTO"
	"boilerplate-go/Services/RabbitMQ"
	subcriberConfig "boilerplate-go/Services/RabbitMQ/subscriberConfig"
)

type subscriber struct {
	Config   subcriberConfig.SubscriberConfig
	Service  RabbitMQ.SubscriberService
	Consumer Subscribers.ControllerConsumer
}

type SubscriberRoutes interface {
	CollectSubscriber()
}

// collection to setup your subscription
func (s subscriber) CollectSubscriber() {
	s.Service.Add(s.Config.ExampleConfig.ConfigsSubsExample, s.Consumer.ExampleConsumer.ExampleSubscriberController)
	s.Service.Add(s.Config.ExampleConfig.ConfigExampleDirectRPC, s.Consumer.ExampleConsumer.ExampleSubscriberRPCController)
	s.Service.Add(s.Config.ExampleConfig.ConfigExampleDirectNonRPC, s.Consumer.ExampleConsumer.ExampleSubscriberNonRPCController)
	s.Service.ExecuteService()
}

// constructor for consumer
func NewConsumer(Service RabbitMQ.SubscriberService, u DTO.Utilities) SubscriberRoutes {
	return subscriber{
		Service: Service,
		Config: subcriberConfig.SubscriberConfig{
			ExampleConfig: subcriberConfig.SubsConfig{},
		},
		Consumer: Subscribers.NewConsumerController(u),
	}
}
