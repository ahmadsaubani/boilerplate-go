package ExampleRepositoryPattern

import "boilerplate-go/Repository"

type exampleRepositoryPatternModules struct {
	repo Repository.Repository
}

type ExampleRepositoryPatternModulesModules interface {
	GetExample() (boolean bool, err error)
}

func NewModules(Repo Repository.Repository) ExampleRepositoryPatternModulesModules {
	return &exampleRepositoryPatternModules{
		repo: Repo,
	}
}
