package RabbitMQ

import (
	subcriberConfig "boilerplate-go/Services/RabbitMQ/subscriberConfig"
	"log"

	"github.com/streadway/amqp"
)

type SubscriberService interface {
	Add(configType subcriberConfig.ConfigType, consumer ConsumerLogic)
	ExecuteService()
}

func (r *rabbit) Add(configType subcriberConfig.ConfigType, consumer ConsumerLogic) {

	if r.connection != nil {
		channel, err := r.connection.Channel()
		if err != nil {
			log.Println(err.Error())
			panic(err)
		}
		consumerConfig := configType()
		Subscriber := &rabbit{
			Config:               &consumerConfig,
			ControllerSubscriber: consumer,
			Subscriber:           nil,
			Channel:              channel,
		}
		if r.Subscriber == nil {
			r.Subscriber = Subscriber
			return
		}
		SubscriberList := r.Subscriber
		for SubscriberList.Subscriber != nil {
			SubscriberList = SubscriberList.Subscriber
		}
		SubscriberList.Subscriber = Subscriber
	}
}

func (r rabbit) ExecuteService() {
	if r.connection != nil {
		list := r.Subscriber
		if r.Subscriber == nil {
			log.Println("No node are available")
			return
		}
		delivery, err := list.Config.InitConfig(list.Channel)
		if err != nil {
			log.Printf("%s", err)
			panic(err.Error())
		}
		go rabbitRun(delivery, list.Channel, list.ControllerSubscriber)
		for list.Subscriber != nil {
			list = list.Subscriber
			delivery, err = list.Config.InitConfig(list.Channel)
			if err != nil {
				log.Printf("%s", err)
				panic(err.Error())
			}
			go rabbitRun(delivery, list.Channel, list.ControllerSubscriber)
		}
		log.Println("ENGINE rabbit consumer start....")
		return
	}
	log.Println("ENGINE rabbit consumer failed to start....")
}

func rabbitRun(delivery <-chan amqp.Delivery, RabbitChannel *amqp.Channel, controllerSubscriber func(data amqp.Delivery, rabbitChannel *amqp.Channel)) {
	for d := range delivery {
		controllerSubscriber(d, RabbitChannel)
	}
}
