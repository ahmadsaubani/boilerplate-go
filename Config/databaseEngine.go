package Config

import (
	"bytes"
	"database/sql"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type connectionPool interface {
	PostgresConnectionPool(Connection interface{})
	RedisConnectionPool(Connection interface{})
}

func (env *database) PostgresConnectionPool(Connection interface{}) {
	Con := Connection.(*sql.DB)
	var buffer bytes.Buffer
	buffer.WriteString("postgres://")
	buffer.WriteString(env.Username + ":" + env.Password)
	buffer.WriteString("@")
	buffer.WriteString(env.Host + ":" + env.Port + "/")
	buffer.WriteString(env.Name)
	buffer.WriteString("?sslmode=disable")
	connection_string := buffer.String()

	Connection, err := sql.Open("postgres", connection_string)
	if err != nil {
		panic(err.Error())
		return
	}
	Connection.(*sql.DB).SetMaxOpenConns(env.Maximum_connection)
	Connection.(*sql.DB).SetConnMaxIdleTime(env.MaximumIdleTime * time.Second)
	*Con = *Connection.(*sql.DB)
	err = Con.Ping()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
		return
	}
	return
}

func (env *database) RedisConnectionPool(Connection interface{}) {
	var buffer bytes.Buffer
	Con := Connection.(*redis.Client)
	buffer.WriteString(env.Host + ":" + env.Port)
	connectionString := buffer.String()

	Connection = redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: env.Password,
		DB:       0,
	})

	*Con = *Connection.(*redis.Client)

	_, err := Con.Ping().Result()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
		return
	}
}
