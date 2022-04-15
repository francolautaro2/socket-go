// this idea is that server listen multiples clientes (cocurrency)
package server

import (
	"fmt"
	"log"
	"net"
)

// Server config
type ServerConfig struct {
	Port *net.TCPAddr
	Addr string
}

// Create Server
func NewServer(port string, Addr string) ServerConfig {
	service, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	var s ServerConfig
	s.Port = service
	s.Addr = Addr
	return s
}

// Listen clients
func (s ServerConfig) RunServer() {
	fmt.Printf("Server is running in %s%s", s.Addr, s.Port)

	ln, err := net.ListenTCP("tcp", s.Port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		conn.Write([]byte("Hello World"))
		conn.Close()
	}
}
