package DTO

import (
	"boilerplate-go/Modules"
	"boilerplate-go/Services/RabbitMQ"
)

type Utilities struct {
	Rabbit          RabbitMQ.ServiceInterface
	RabbitRepublish RabbitMQ.ServiceInterface
	Modules         Modules.Modules
}
