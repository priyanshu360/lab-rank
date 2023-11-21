package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/priyanshu360/lab-rank/dashboard/api"
	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/utils"
)

func main() {
	if err := godotenv.Load("local.env"); err != nil {
		slog.Warn("Error in loading env file, Generate .env file")
	}

	dbConf := config.NewDBConfig()
	loggerConf := config.InitLoggerConfig()
	serverConf := config.NewServerConfig()

	api.InitDB(dbConf)
	api.StartHttpServer(serverConf)

	logger := utils.NewLogger(loggerConf)
	slog.SetDefault(logger)
}
