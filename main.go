package main

import (
	"log"

	"github.com/blacdev/werant/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
