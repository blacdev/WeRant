package main

import (
	"github.com/blacdev/werant/config"
	"github.com/blacdev/werant/controller"
	"github.com/blacdev/werant/server"
	"github.com/blacdev/werant/service"
	"log"
)

func main() {
	cts := controller.NewContainer()
	sc := service.NewContainer()

	//todo: handle interrupt (Ctrl+C)

	err := server.Start(config.GetServerAddress(), cts, sc)
	if err != nil {
		log.Fatal(err)
	}
}
