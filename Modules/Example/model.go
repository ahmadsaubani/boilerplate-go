package Example

import (
	"boilerplate-go/Modules/ModuleDTO"
	"boilerplate-go/Repository"
	"boilerplate-go/Services/RabbitMQ"
	"database/sql"
)

type example struct {
	repo            Repository.Repository
	rabbit          RabbitMQ.ServiceInterface
	rabbitRepublish RabbitMQ.ServiceInterface
}

type ExampleModules interface {
	DBStatus() sql.DBStats
	IsExists(yourName string) (res bool)
}

func NewExampleModules(moduleConfig ModuleDTO.ModulesConfig) ExampleModules {
	return &example{
		repo:            moduleConfig.Repo,
		rabbit:          moduleConfig.RabbitPublish,
		rabbitRepublish: moduleConfig.RabbitRepublish,
	}
}
