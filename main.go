package main

import (
	"fmt"
	"log/slog"
	"os"
	"web/configs"
	_ "web/docs"
	"web/internal/comment"
	"web/internal/pkg/db"
	"web/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
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

func initializeDB(endpoint string) *gorm.DB {
	database := db.Connect(endpoint)
	if err := database.AutoMigrate(&user.User{}); err != nil {
		slog.Error("db migrate error", slog.Any("err", err))
		os.Exit(1)
	}
	if err := database.AutoMigrate(&comment.Comment{}); err != nil {
		slog.Error("db migrate error", slog.Any("err", err))
		os.Exit(1)
	}

	return database
}

// @title       REST API
// @version     1.0
// @description API documentation
func main() {
	config := retrieveConfig()

	db := initializeDB(config.Database.Endpoint)

	userRepo := user.NewDBRepository(db)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.RegisterRoutes(r, userRepo)
	comment.RegisterRoutes(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
