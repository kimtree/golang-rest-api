package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"web/api"
	"web/configs"
	"web/internal/comment"
	"web/internal/user"
	"web/pkg/db"
)

func retrieveConfig() configs.Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("failed to connect DB", slog.Any("err", err))
		os.Exit(1)
	}

	var config configs.Config
	if err := viper.Unmarshal(&config); err != nil {
		slog.Error("failed to connect DB")
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("config: %+v", config))

	return config
}

func initializeDB(endpoint string) {
	db.Connect(endpoint)
	if err := db.Migrate(&user.User{}); err != nil {
		slog.Error("db migrate error", slog.Any("err", err))
		os.Exit(1)
	}
	if err := db.Migrate(&comment.Comment{}); err != nil {
		slog.Error("db migrate error", slog.Any("err", err))
		os.Exit(1)
	}
}

func main() {
	config := retrieveConfig()

	initializeDB(config.Database.Endpoint)

	r := gin.Default()
	r.Use(gin.Recovery())
	api.RegisterRoutes(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
