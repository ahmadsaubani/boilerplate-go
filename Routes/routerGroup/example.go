package routerGroup

import (
	"boilerplate-go/Controller/Api"

	"github.com/gin-gonic/gin"
)

type iExampleGroup interface {
	ExampleApiGroup(group *gin.RouterGroup, api Api.ControllerExample)
}

func (r exampleGroup) ExampleApiGroup(group *gin.RouterGroup, api Api.ControllerExample) {
	group.GET("/ping", api.ExamplePingController)
	group.POST("/rpc", api.ExamplePublishToRabbit)
	group.POST("/non-rpc", api.ExamplePublishToRabbitNonRPC)
}

type exampleGroup struct{}

func newExampleGroup() iExampleGroup {
	return &exampleGroup{}
}
