package main

import (
	"broker/config"
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	rabbit := config.GetConnect()
	defer rabbit.Close()

	ch := config.OpenCh(rabbit)
	defer ch.Close()

	queue := config.Declare(ch)

	msg := "Ini adalah message yang akan di kirim"
	er := ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if er != nil {
		fmt.Println("Failed Publish a Message")
		panic(er)
	}

}
