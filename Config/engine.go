package Config

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
)

const EnvirontmentPath = "../Environtment/env.yml"
const Postgres = "postgres"
const Redis = "redis"

func GetEnvironment(env string) Config {
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), EnvirontmentPath)
	_, err := os.Stat(envPath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
		return nil
	}
	content, err := ioutil.ReadFile(envPath)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	var config envFile = content
	return config
}

func (e envFile) LoadConfig() *ConfigSetting {

	var config Environment

	err := yaml.Unmarshal([]byte(string(e)), &config)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	if config.App.Debug == false {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("Environment Config load successfully!")
	return &ConfigSetting{&config, nil, &config.App, &config}
}

func (e *Environment) BuildConnection() DbConInterface {
	var connectionPool connectionPool = &database{}
	var dbCon DbCon
	for i := 0; i < len(e.Databases); i++ {
		connectionPool = &e.Databases[i]
		switch e.Databases[i].Engine {
		case Postgres:
			con := sql.DB{}
			log.Println("ENGINE " + Postgres + " start....")
			connectionPool.PostgresConnectionPool(&con)
			dbCon.setSqlConnection(DbSqlConfigName(e.Databases[i].Connection), &con)
		case Redis:
			con := redis.Client{}
			log.Println("ENGINE " + Redis + " start....")
			connectionPool.RedisConnectionPool(&con)
			dbCon.setRedisConnection(dbRedisConfigName(e.Databases[i].Connection), &con)
		}

	}

	return &dbCon
}

func (app *app) Run(route *gin.Engine) {
	app.runWithHttp(route)
}

func (app *app) runWithHttp(route *gin.Engine) {
	log.Println("Http Service running ....")
	address := app.Host + ":" + app.Port
	err := route.Run(address)
	if err != nil {
		//panic error
		panic(err)
	}
}
