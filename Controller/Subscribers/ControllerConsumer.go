package Subscribers

import (
	"boilerplate-go/Controller/Subscribers/ExampleSubscriberController"
	"boilerplate-go/DTO"
)

type ControllerConsumer struct {
	ExampleConsumer ExampleSubscriberController.ExampleConsumerInterface
}

// constructure
func NewConsumerController(u DTO.Utilities) ControllerConsumer {
	return ControllerConsumer{
		ExampleConsumer: ExampleSubscriberController.NewConsumer(u),
	}
}
