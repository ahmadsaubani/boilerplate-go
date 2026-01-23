package subcriberConfig

const (
	ExampleTopic         = "example.topic"
	ExampleQueueName     = "example.example-config"
	ExampleQueueConsumer = "example.example-consumer"
	ExampleBindingKey    = "example.binding-key"
)

type exampleConfig interface {
	ConfigsSubsExample() SubsConfig
}

func (s SubsConfig) ConfigsSubsExample() SubsConfig {
	s.ConfigExchange.ExchangeName = ExampleTopic
	s.ConfigQueue.Name = ExampleQueueName
	s.ConfigSubscribe.Subscriber = ExampleQueueConsumer
	s.ConfigSubscribe.AutoAck = true
	s.setDefaultTopicConfig(ExampleBindingKey)
	return s
}
