package RabbitMQ

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

const ContentTypeJSON = "application/json"

func (c *RabbitConfig) InitConnectionSubscriber(ApplicationName string) {
	var err error
	var buffer bytes.Buffer

	buffer.WriteString(c.Config.Scheme + "://")
	buffer.WriteString(c.Config.Username + ":" + c.Config.Password)
	buffer.WriteString("@")
	buffer.WriteString(c.Config.Host + ":" + c.Config.Port)
	buffer.WriteString(c.Config.Vhost)
	connectionString := buffer.String()

	RabbitConfig := amqp.Config{
		Properties: amqp.Table{
			"connection_name": "Subscriber" + ApplicationName,
		},
	}
	if c.Config.Host != "" {
		c.connectionSubscriber, err = amqp.DialConfig(connectionString, RabbitConfig)
		if err != nil {
			panic(err.Error())
			return
		}
		log.Println("ENGINE rabbit start....")
		return
	}
	log.Println("ENGINE rabbit not found....")
}

func (c *RabbitConfig) InitConnectionPublisher(ApplicationName string) {
	var err error
	var buffer bytes.Buffer

	buffer.WriteString(c.Config.Scheme + "://")
	buffer.WriteString(c.Config.Username + ":" + c.Config.Password)
	buffer.WriteString("@")
	buffer.WriteString(c.Config.Host + ":" + c.Config.Port)
	buffer.WriteString(c.Config.Vhost)
	connectionString := buffer.String()

	RabbitConfig := amqp.Config{
		Properties: amqp.Table{
			"connection_name": "Publisher" + ApplicationName,
		},
	}
	if c.Config.Host != "" {
		c.connectionRepublish, err = amqp.DialConfig(connectionString, RabbitConfig)
		if err != nil {
			panic(err.Error())
			return
		}
		log.Println("ENGINE rabbit start....")
		return
	}
	log.Println("ENGINE rabbit not found....")
}

type RabbitKey string
type RabbitExhangeName string

func (s service) Publish(key RabbitKey, exhange RabbitExhangeName, data interface{}) error {
	ch, err := s.Con.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	DataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = ch.Publish(string(exhange), string(key), false, false, amqp.Publishing{
		ContentType: ContentTypeJSON,
		Body:        DataByte,
	})
	return err
}

func (s service) PublishRpc(key RabbitKey, exhange RabbitExhangeName, data interface{}) (msg []byte, err error) {
	ch, err := s.Con.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()
	DataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	corrId := uuid.NewV4().String()
	queue, err := ch.QueueDeclare(
		s.queueTemp(key),
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	messages, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	err = ch.Publish(string(exhange), string(key), false, false, amqp.Publishing{
		ContentType:   ContentTypeJSON,
		Body:          DataByte,
		ReplyTo:       queue.Name,
		CorrelationId: corrId,
	})
	if err != nil {
		return nil, err
	}
	for msg := range messages {
		if corrId == msg.CorrelationId {
			return msg.Body, nil
		}
	}

	return nil, err
}

func (s service) ReplyRpc(delivery amqp.Delivery, rabbitChannel *amqp.Channel, data interface{}) (err error) {
	DataByte, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = rabbitChannel.Publish("", delivery.ReplyTo, false, false, amqp.Publishing{
		ContentType:   ContentTypeJSON,
		Body:          DataByte,
		CorrelationId: delivery.CorrelationId,
	})
	if err != nil {
		delivery.Ack(false)
		return
	}
	delivery.Ack(false)
	return
}

func (s service) queueTemp(prefix RabbitKey) string {
	uuid := uuid.NewV4().String()
	return fmt.Sprintf("%s-%s", string(prefix), uuid)
}
