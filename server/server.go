package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Server ...
type Server struct {
	host    string
	port    string
	clients []*Client
}

// Client ...
type Client struct {
	conn   net.Conn
	server *Server
}

// Config ...
type Config struct {
	Host string
	Port string
}

// New ...

func NewServer(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

// Run ...
func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn:   conn,
			server: server,
		}
		server.clients = append(server.clients, client)
		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	reader := bufio.NewReader(client.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			client.conn.Close()
			return
		}
		for _, client1 := range client.server.clients {
			client1.conn.Write([]byte(message))
		}
	}
}
