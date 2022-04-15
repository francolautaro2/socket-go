package main

import (
	"os"

	"github.com/francolautaro2/socket-go/server"
)

func main() {
	// port from user
	p := os.Args[1]

	//Create server config
	s := server.NewServer(p, "127.0.0.1")
	s.RunServer() // Run server
}
