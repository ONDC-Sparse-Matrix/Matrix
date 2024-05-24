package utils

import (
	"centralCDN/pkg/types"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectQueue() *amqp.Connection{
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	return conn
}


func getPubChannel(Connection *amqp.Connection,channelName string) *amqp.Channel {
	ch, err := Connection.Channel()
	if err != nil {
		panic(err)
	}

	ch.QueueDeclare(channelName, true, false, false, false, nil)
	return ch
}

func getSubChannel(Connection *amqp.Connection,channelName string) *amqp.Channel {
	ch, err := Connection.Channel()
	if err != nil {
		panic(err)
	}

	ch.QueueDeclare(channelName, true, false, false, false, nil)
	return ch
}

func PublishCache(Connection *amqp.Connection,clientId string,cache []types.PincodeInfo)error{
	ch := getPubChannel(Connection,clientId)
	defer ch.Close()
	body, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	ch.Publish("", clientId, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	return nil
}

func ConsumeCache(Connection *amqp.Connection,clientId string) (<-chan amqp.Delivery,error){
	ch := getSubChannel(Connection,clientId)
	defer ch.Close()
	msgs, err := ch.Consume(clientId, "", true, false, false, false, nil)
	if err != nil {
		return nil,err
	}
	return msgs,nil
}

func PublishSeeding(Connection *amqp.Connection,data []byte)error{
	ch := getPubChannel(Connection,"merchant_data")
	defer ch.Close()
	ch.Publish("", "merchant_data", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	})
	return nil
}

func ConsumeSeeding(Connection *amqp.Connection) (<-chan amqp.Delivery,error){
	ch := getSubChannel(Connection,"merchant_data")
	defer ch.Close()
	msgs, err := ch.Consume("merchant_data", "", true, false, false, false, nil)
	if err != nil {
		return nil,err
	}
	return msgs,nil
}
