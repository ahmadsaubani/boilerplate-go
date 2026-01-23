package subcriberConfig

import "github.com/streadway/amqp"

func (s SubsConfig) InitConfig(rabbitChannel *amqp.Channel) (delivery <-chan amqp.Delivery, err error) {
	// set up your config for specific module here!
	if s.ConfigExchange.Status == true {
		if err = rabbitChannel.ExchangeDeclare(
			s.ConfigExchange.ExchangeName,
			s.ConfigExchange.ExchangeType,
			s.ConfigExchange.Durable,
			s.ConfigExchange.AutoDelete,
			s.ConfigExchange.Internal,
			s.ConfigExchange.NoWait,
			s.ConfigExchange.Args,
		); err != nil {
			return
		}
	}

	if _, err := rabbitChannel.QueueDeclare(
		s.ConfigQueue.Name,
		s.ConfigQueue.Durable,
		s.ConfigQueue.AutoDelete,
		s.ConfigQueue.Exclusive,
		s.ConfigQueue.NoWait,
		s.ConfigQueue.Args,
	); err != nil {
		return nil, err
	}

	if s.ConfigBinding != nil {
		for _, values := range s.ConfigBinding {
			if err = rabbitChannel.QueueBind(
				values.Queue,
				values.Key,
				values.ExchangeName,
				values.NoWait,
				values.Args,
			); err != nil {
				return
			}
		}
	}

	if delivery, err = rabbitChannel.Consume(
		s.ConfigSubscribe.Queue,
		s.ConfigSubscribe.Subscriber,
		s.ConfigSubscribe.AutoAck,
		s.ConfigSubscribe.Exclusive,
		s.ConfigSubscribe.NoLocal,
		s.ConfigSubscribe.NoWait,
		s.ConfigSubscribe.Args,
	); err != nil {
		return
	}

	return
}

func (config *SubsConfig) setDefaultDirectConfig(bindingKeys ...string) {

	/*
		what need to setup :
			1.configExchange.ExchangeName
			2.configQueue.Name
			3.configSubscribe.Subscriber
	*/

	config.ConfigExchange.Status = true
	config.ConfigExchange.ExchangeType = ExchangeDirect.Type()
	config.ConfigExchange.Durable = true

	config.ConfigQueue.Durable = true

	for _, v := range bindingKeys {
		config.ConfigBinding = append(config.ConfigBinding,
			configBinding{
				Status:       true,
				Queue:        config.ConfigQueue.Name,
				Key:          v,
				ExchangeName: config.ConfigExchange.ExchangeName,
			})
	}

	config.ConfigSubscribe.Queue = config.ConfigQueue.Name
}

func (config *SubsConfig) setDefaultTopicConfig(bindingKeys ...string) {

	/*
		what need to setup :
			1.configExchange.ExchangeName
			2.configQueue.Name
			3.configSubscribe.Subscriber
	*/

	config.ConfigExchange.Status = true
	config.ConfigExchange.ExchangeType = ExchangeTopic.Type()
	config.ConfigExchange.Durable = true

	config.ConfigQueue.Durable = true

	for _, v := range bindingKeys {
		config.ConfigBinding = append(config.ConfigBinding,
			configBinding{
				Status:       true,
				Queue:        config.ConfigQueue.Name,
				Key:          v,
				ExchangeName: config.ConfigExchange.ExchangeName,
			})
	}

	config.ConfigSubscribe.Queue = config.ConfigQueue.Name
}

func (config *SubsConfig) setDefaultWithoutExchangeConfig(bindingKeys ...string) {

	/*
		what need to setup :
			1.configExchange.ExchangeName
			2.configQueue.Name
			3.configSubscribe.Subscriber
	*/

	config.ConfigExchange.Status = false
	config.ConfigExchange.Durable = true

	config.ConfigQueue.Durable = true

	for _, v := range bindingKeys {
		config.ConfigBinding = append(config.ConfigBinding,
			configBinding{
				Status:       true,
				Queue:        config.ConfigQueue.Name,
				Key:          v,
				ExchangeName: config.ConfigExchange.ExchangeName,
			})
	}

	config.ConfigSubscribe.Queue = config.ConfigQueue.Name
}
