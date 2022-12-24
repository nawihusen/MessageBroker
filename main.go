package main

import (
	"broker/config"
	"fmt"
)

func main() {
	rabbit := config.GetConnect()
	defer rabbit.Close()

	ch, er := rabbit.Channel()
	if er != nil {
		fmt.Println("Failet To Open Chanel")
	}
	defer ch.Close()

}
