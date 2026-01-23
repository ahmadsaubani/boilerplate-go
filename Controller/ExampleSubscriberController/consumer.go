package ExampleSubscriberController

import (
	"boilerplate-go/DTO"
	"boilerplate-go/Library/Helper"
	"encoding/json"
	"net/http"

	"github.com/streadway/amqp"
)

type ExampleConsumerInterface interface {
	ExampleSubscriberController(data amqp.Delivery, rabbitChannel *amqp.Channel)
}

func NewConsumer(o DTO.Utilities) ExampleConsumerInterface {
	return &exampleConsumer{o}
}

func (o exampleConsumer) ExampleSubscriberController(data amqp.Delivery, rabbitChannel *amqp.Channel) {

	var params DTO.Utilities
	err := json.Unmarshal(data.Body, &params)
	if err != nil {
		dataResponse := Helper.SetResponsePublisher(err.Error(), nil, 400)
		err = o.Rabbit.ReplyRpc(data, rabbitChannel, dataResponse)
		if err != nil {
			return
		}
		return
	}

	msg, err := o.RabbitRepublish.PublishRpc("bindingKey",
		"exchangeName", "data")
	if err != nil {
		return
	}

	respRPC := Helper.GetResponseFromRPCReply(msg)
	if respRPC.Code != http.StatusOK {
		return
	}

	dataResponse := Helper.SetResponsePublisher("", http.StatusText(http.StatusOK), http.StatusOK)
	err = o.Rabbit.ReplyRpc(data, rabbitChannel, dataResponse)
	if err != nil {
		return
	}
	return
}
