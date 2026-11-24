package ExampleSubscriberController

import (
	"boilerplate-go/DTO"
	"boilerplate-go/Library/Helper"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/streadway/amqp"
)

type ExampleConsumerInterface interface {
	ExampleSubscriberController(data amqp.Delivery, rabbitChannel *amqp.Channel)
	ExampleSubscriberRPCController(data amqp.Delivery, rabbitChannel *amqp.Channel)
	ExampleSubscriberNonRPCController(data amqp.Delivery, rabbitChannel *amqp.Channel)
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

func (o exampleConsumer) ExampleSubscriberRPCController(data amqp.Delivery, rabbitChannel *amqp.Channel) {
	var params DTO.ExampleRPC
	err := json.Unmarshal(data.Body, &params)
	if err != nil {
		dataResponse := Helper.SetResponsePublisher(err.Error(), nil, 400)
		err = o.Rabbit.ReplyRpc(data, rabbitChannel, dataResponse)
		if err != nil {
			return
		}
		return
	}

	resp := make(map[string]interface{})
	resp["your_name"] = params.YourName

	dataResponse := Helper.SetResponsePublisher("", resp, http.StatusOK)
	err = o.Rabbit.ReplyRpc(data, rabbitChannel, dataResponse)
	if err != nil {
		return
	}
	return
}

func (o exampleConsumer) ExampleSubscriberNonRPCController(data amqp.Delivery, rabbitChannel *amqp.Channel) {
	var params DTO.ExampleRPC
	err := json.Unmarshal(data.Body, &params)
	if err != nil {
		return
	}
	exampleModule := o.Modules.ExampleModule
	exists := exampleModule.IsExists(params.YourName)

	if !exists {
		fmt.Println("No data provided ")
		return
	}
	fmt.Println("data provided : ", params.YourName)
	return
}
