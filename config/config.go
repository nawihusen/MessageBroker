package config

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

type QueueConfig struct {
	Username string
	Password string
	Host     string
	Port     string
}

func GetConnect() *amqp.Connection {
	con, er := Rabbit()
	if er != nil {
		fmt.Println("Failed To Connect")
	}

	return con
}

func Rabbit() (*amqp.Connection, error) {
	pattern := "amqp://%s:%s@%s:%s/"

	clientUrl := fmt.Sprintf(pattern,
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
	)
	conn, er := amqp.Dial(clientUrl)
	if er != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(er)
	}

	fmt.Println("Success Get Connect To Rabbit")
	return conn, nil
}

func OpenCh(con *amqp.Connection) *amqp.Channel {
	ch, er := con.Channel()
	if er != nil {
		fmt.Println("Failet To Open Chanel")
		panic(er)
	}

	return ch
}

func Declare(ch *amqp.Channel) amqp.Queue {
	queue, er := ch.QueueDeclare("Hello", false, false, false, false, nil)
	if er != nil {
		fmt.Println("Failed to declare a queue")
		panic(er)
	}

	return queue
}
