package main

import (
	"fmt"
	"github.com/lxflp/tcp-chat/server"
)

func main() {
	var cfg server.Config
	cfg.Port = "8090"
	cfg.Host = "0.0.0.0"
	var srv *server.Server
	srv = server.NewServer(&cfg)
	fmt.Println("Сервер запущен")
	srv.Run()
}
