package main

import (
	"log"

	"tvubbs/server"
)

func init() {
	// do initialization stuff here.. maybe
}

func main() {
	log.Println("Starting telnet chat server")

	s, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to initialize TCP listener: %s", err.Error())
	}

	s.Serve()
}
