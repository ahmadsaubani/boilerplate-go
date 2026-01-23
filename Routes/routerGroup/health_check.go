package routerGroup

import (
	"boilerplate-go/Library/Helper"

	"github.com/gin-gonic/gin"
)

type iHealthCheckGroup interface {
	HealthCheckApiGroup(group *gin.RouterGroup)
}

func (r healthCheckGroup) HealthCheckApiGroup(group *gin.RouterGroup) {
	group.POST("/ping", func(c *gin.Context) {
		Helper.HttpResponseSuccess(c, map[string]string{
			"message": "pong",
		})
	})
}

type healthCheckGroup struct{}

func newHealthCheckGroup() iHealthCheckGroup {
	return &healthCheckGroup{}
}
