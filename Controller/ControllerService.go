package Controller

import (
	"boilerplate-go/Controller/Api"
	"boilerplate-go/DTO"
)

type Controller struct {
	Example Api.ControllerExample
}

func InitControllerService(u DTO.Utilities) Controller {
	return Controller{
		Example: Api.NewControllerExample(u),
	}
}
