package ExampleRepositoryPattern

import (
	"boilerplate-go/Config"
)

type Repository interface {
	TestRepo() (data bool, err error)
}

type exampleRepositoryPattern struct {
	dbCon Config.DbConInterface
}

func RepositoryNew(dbCon Config.DbConInterface) Repository {
	return &exampleRepositoryPattern{
		dbCon: dbCon,
	}
}
