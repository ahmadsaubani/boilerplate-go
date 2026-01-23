package subcriberConfig

import "github.com/streadway/amqp"

type SubsConfig struct {
	ConfigExchange  configExchange
	ConfigQueue     configQueue
	ConfigBinding   []configBinding
	ConfigSubscribe configSubscribe
}

type configExchange struct {
	Status       bool
	ExchangeName string
	ExchangeType string
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Args         amqp.Table
}

type configQueue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type configBinding struct {
	Status       bool
	Queue        string
	Key          string
	ExchangeName string
	NoWait       bool
	Args         amqp.Table
}

type configSubscribe struct {
	Queue      string
	Subscriber string
	AutoAck    bool
	Exclusive  bool
	NoLocal    bool
	NoWait     bool
	Args       amqp.Table
}

// type for rabbit config
type ConfigType func() SubsConfig

// constructure for Rabbit config
type SubscriberConfig struct {
	ExampleConfig exampleConfig
}

type Exchanges int

const (
	ExchangeDirect Exchanges = 1
	ExchangeFanout Exchanges = 2
	ExchangeHeader Exchanges = 3
	ExchangeTopic  Exchanges = 4
	ExchangeEnd    Exchanges = 5
)

func (e Exchanges) Name() string {
	exchangeType := [...]string{
		"amq.direct",
		"amq.fanout",
		"amq.headers",
		"amq.topic",
	}
	return exchangeType[e-1]
}

func (e Exchanges) Type() string {
	exchangeType := [...]string{
		amqp.ExchangeDirect,
		amqp.ExchangeFanout,
		amqp.ExchangeHeaders,
		amqp.ExchangeTopic,
	}
	return exchangeType[e-1]
}

func (e Exchanges) IsValid() bool {
	if int(e) >= int(ExchangeEnd) || int(e) <= 0 {
		return false
	}
	return true
}
