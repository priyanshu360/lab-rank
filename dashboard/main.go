package main

import (
	"fmt"
	"log/slog"

	"github.com/priyanshu360/lab-rank/dashboard/api"
	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading env file, Generate .env file")
		return
	}
	loggerConf := config.InitLoggerConfig()
	logger := utils.NewLogger(loggerConf)
	slog.SetDefault(logger)
	api.InitDB()
	config := config.NewEnvServerConfig()
	api.StartHttpServer(config)
}
