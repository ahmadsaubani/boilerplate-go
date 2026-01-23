package ExampleRepositoryPattern

func (d exampleRepositoryPattern) TestRepo() (data bool, err error) {
	connection := d.dbCon.PostgreMainCon()
	defer connection.Close()

	data = true
	return
}
