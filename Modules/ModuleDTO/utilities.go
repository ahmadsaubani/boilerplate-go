package ModuleDTO

import (
	"boilerplate-go/Repository"
	"boilerplate-go/Services/RabbitMQ"
)

type ModulesConfig struct {
	Repo            Repository.Repository
	RabbitPublish   RabbitMQ.ServiceInterface
	RabbitRepublish RabbitMQ.ServiceInterface
}
