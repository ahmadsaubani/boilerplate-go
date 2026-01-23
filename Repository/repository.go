package Repository

import (
	"boilerplate-go/Config"
	"boilerplate-go/Repository/Example"
	"boilerplate-go/Repository/ExampleRepositoryPattern"
)

// CONSTRUCTOR STRUCT FOR ALL REPOSITORY
type Repository struct {
	Example     Example.Repository
	ExampleRepo ExampleRepositoryPattern.Repository
}

// REPOSITORY INITIALIZATION
func InitRepo(dbCon Config.DbConInterface) Repository {
	return Repository{
		Example:     Example.RepositoryNew(dbCon),
		ExampleRepo: ExampleRepositoryPattern.RepositoryNew(dbCon),
	}
}
