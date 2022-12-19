package main

import (
	"log"

	"tvubbs/bbsconfig"
	"tvubbs/server"
)

func init() {
	// do initialization stuff here.. maybe
}

func main() {
	log.Println("Starting BBS...")
	if err := bbsconfig.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %s", err.Error())
	}
	s, err := server.NewServer()
	if err != nil {
		log.Fatalf("Failed to initialize TCP listener: %s", err.Error())
	}

	s.Serve()
}
