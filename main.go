package main

import (
	"log"

	"tvubbs/bbsconfig"
	"tvubbs/server"
)

func init() {
	bbsconfig.LoadConfig()
}

func main() {
	log.Println("Starting telnet chat server")

	s, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to initialize TCP listener: %s", err.Error())
	}

	s.Serve()
}
