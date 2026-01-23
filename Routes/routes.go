package Routes

import (
	"boilerplate-go/Controller"
	"boilerplate-go/Library/Helper"
	"boilerplate-go/Routes/routerGroup"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	Controller Controller.Controller
	Gin        *gin.Engine
}

func (app *Routes) CollectRoutes() *gin.Engine {

	appRoute := app.Gin
	apiGroup := routerGroup.RoutesGroupCollection

	appRoute.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			Helper.HttpResponseError(c, http.StatusInternalServerError, err)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	healthCheck := appRoute.Group("/")

	apiGroup.HealtCheck.HealthCheckApiGroup(healthCheck)
	return appRoute
}
