package main

import (
	"broker/config"
	"fmt"
)

func main() {
	rabbit := config.GetConnect()
	defer rabbit.Close()

	ch := config.OpenCh(rabbit)
	defer ch.Close()

	queue := config.Declare(ch)

	msg, er := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if er != nil {
		fmt.Println("Failed to register a consumer")
		panic(er)
	}

	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Printf("Ini adalah pesan diterima : %s", d.Body)
		}
	}()

	fmt.Println("Success All Procces")
	fmt.Println("Waiting For A Message")
	<-forever

	// func ini akan terus berjalan karena channel yang di gunakan tidak pernah di tangkap
	// sekali
}
