package subcriberConfig

import "boilerplate-go/Constant"

const (
	ExampleTopic         = "example.topic"
	ExampleQueueName     = "example.example-config"
	ExampleQueueConsumer = "example.example-consumer"
	ExampleBindingKey    = "example.binding-key"
)

type exampleConfig interface {
	ConfigsSubsExample() SubsConfig
	ConfigExampleDirectRPC() SubsConfig
	ConfigExampleDirectNonRPC() SubsConfig
}

func (s SubsConfig) ConfigsSubsExample() SubsConfig {
	s.ConfigExchange.ExchangeName = ExampleTopic
	s.ConfigQueue.Name = ExampleQueueName
	s.ConfigSubscribe.Subscriber = ExampleQueueConsumer
	s.ConfigSubscribe.AutoAck = true
	s.setDefaultTopicConfig(ExampleBindingKey)
	return s
}

func (s SubsConfig) ConfigExampleDirectRPC() SubsConfig {
	s.ConfigExchange.ExchangeName = Constant.ExampleExchange
	s.ConfigQueue.Name = Constant.ExampleQueueRPCName
	s.ConfigSubscribe.Subscriber = Constant.ExampleRPCConsumer
	s.ConfigSubscribe.AutoAck = false
	s.setDefaultDirectConfig(Constant.ExampleRabbitKey)
	return s
}

func (s SubsConfig) ConfigExampleDirectNonRPC() SubsConfig {
	s.ConfigExchange.ExchangeName = Constant.ExampleExchange
	s.ConfigQueue.Name = Constant.ExampleNonRPCQueueName
	s.ConfigSubscribe.Subscriber = Constant.ExampleNonRPCRPCConsumer
	s.ConfigSubscribe.AutoAck = true
	s.setDefaultDirectConfig(Constant.ExampleNonRPCRabbitKey)
	return s
}
