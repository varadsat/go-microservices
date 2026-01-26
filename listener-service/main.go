package main

import (
	"listener-service/event"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//connect to rabbitmq
	rabitConn, err := connect()

	if err != nil {
		log.Fatalln(err)
	}
	defer rabitConn.Close()
	log.Println("Connected to RabbitMQ!")

	//start listening to messages
	log.Println("Listening for and consuming RabbitMQ messages...")

	//create consumer
	consumer, err := event.NewConsumer(rabitConn)
	if err != nil {
		panic(err)
	}

	//watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}
func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		con, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			log.Println("RabbitMQ not ready yet...")
			counts++
		} else {
			connection = con
			break
		}
		if counts > 5 {
			log.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
	}
	return connection, nil
}
