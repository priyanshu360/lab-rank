package main

import (
	"log/slog"

	"github.com/priyanshu360/lab-rank/dashboard/api"
	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/utils"
)

func main() {
	loggerConf := config.InitLoggerConfig()
	logger := utils.NewLogger(loggerConf)
	slog.SetDefault(logger)
	api.InitDB()
	config := config.NewEnvServerConfig()
	api.StartHttpServer(config)
}
