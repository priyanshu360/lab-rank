package main

import (
	"github.com/priyanshu360/lab-rank/dashboard/api"
	"github.com/priyanshu360/lab-rank/dashboard/config"
)

func main() {
	api.InitDB()
	config := config.NewEnvServerConfig()
	api.StartHttpServer(config)
}
