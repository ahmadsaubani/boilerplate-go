package ExampleRepositoryPattern

func (u exampleRepositoryPatternModules) GetExample() (boolean bool, err error) {
	boolean, err = u.repo.ExampleRepo.TestRepo()

	if err != nil {
		return
	}
	return
}
