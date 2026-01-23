package Config

import (
	"time"

	"github.com/gin-gonic/gin"
)

type app struct {
	Appname     string `yaml:"name"`
	Debug       bool   `yaml:"debug"`
	Port        string `yaml:"port"`
	Service     string `yaml:"service"`
	Certificate string `yaml:"certificate"`
	Pem_key     string `yaml:"pem_key"`
	Host        string `yaml:"host"`
}

type database struct {
	Name               string        `yaml:"name"`
	Username           string        `yaml:"username"`
	Password           string        `yaml:"password"`
	Port               string        `yaml:"port"`
	Engine             string        `yaml:"engine"`
	Host               string        `yaml:"host"`
	Maximum_connection int           `yaml:"maximum_connection"`
	MaximumIdleTime    time.Duration `yaml:"maximum_idle_time"`
	Usage              string        `yaml:"usage"`
	Connection         string        `yaml:"connection"`
}
type dbConfig []database

type Environment struct {
	App       app          `yaml:"app"`
	Databases dbConfig     `yaml:"databases"`
	RabbitMQ  RabbitConfig `yaml:"rabbitmq"`
}

type RabbitConfig struct {
	Scheme   string `yaml:"scheme"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Vhost    string `yaml:"vhost"`
}
type envFile []byte
type Config interface {
	LoadConfig() *ConfigSetting
}
type Db interface {
	BuildConnection() DbConInterface
}

type ConfigSetting struct {
	Database    Db
	Routes      RouteInterface
	HttpEngine  HttpEngine
	Environment *Environment
}

type HttpEngine interface {
	runWithHttp(route *gin.Engine)
	Run(route *gin.Engine)
}

type RouteInterface interface {
	CollectRoutes() *gin.Engine
}
