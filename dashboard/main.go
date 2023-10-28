package main

import (
	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/server"
)

func main() {
	server.InitDB()
	config := config.NewEnvServerConfig()
	server.StartHttpServer(config)
}
