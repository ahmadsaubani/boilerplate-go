package Services

import (
	"boilerplate-go/Config"
	"boilerplate-go/Controller"
	"boilerplate-go/DTO"
	"boilerplate-go/Modules"
	"boilerplate-go/Modules/Example"
	"boilerplate-go/Modules/ExampleRepositoryPattern"
	"boilerplate-go/Modules/ModuleDTO"
	"boilerplate-go/Repository"
	"boilerplate-go/Routes"
	"boilerplate-go/Services/RabbitMQ"
	"flag"

	"github.com/gin-gonic/gin"
)

var AppEnv = flag.String("env", "", "define environment")
var routesConfig *gin.Engine

func init() {
	flag.Parse()
	gin.SetMode(gin.DebugMode)
	routesConfig = gin.Default()
	if *AppEnv == "" {
		*AppEnv = "env"
	}
}

// AppInitialization Application Engine initialization
func AppInitialization() {
	newConfig := Config.GetEnvironment(*AppEnv).LoadConfig()
	connection := newConfig.Database.BuildConnection()
	service := serviceInit(newConfig.Environment)

	RabbitService := service.RabbitMQ
	RabbitService.InitConnectionPublisher(newConfig.Environment.App.Appname)
	RabbitService.InitConnectionSubscriber(newConfig.Environment.App.Appname)

	initRepo := Repository.InitRepo(connection)
	ModulesConfig := ModuleDTO.ModulesConfig{
		Repository.InitRepo(connection),
		service.RabbitMQ.GetServicesSubscriber(),
		service.RabbitMQ.GetServices(),
	}

	utilities := DTO.Utilities{
		Rabbit:          service.RabbitMQ.GetServicesSubscriber(),
		RabbitRepublish: service.RabbitMQ.GetServices(),
		Modules: Modules.Modules{
			ExampleModule:     Example.NewExampleModules(ModulesConfig),
			ExampleRepoModule: ExampleRepositoryPattern.NewModules(initRepo),
		},
	}

	newConfig.Routes = &Routes.Routes{
		Gin:        routesConfig,
		Controller: Controller.InitControllerApi(utilities),
	}
	routes := newConfig.Routes.CollectRoutes()
	newConfig.HttpEngine.Run(routes)

}

// Initialization Application Services and Dependency
func serviceInit(Env *Config.Environment) service {
	serv := service{
		RabbitMQ: &RabbitMQ.RabbitConfig{Config: &Env.RabbitMQ},
	}
	return serv
}
