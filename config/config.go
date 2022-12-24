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
