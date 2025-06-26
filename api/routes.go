package api

import (
	"github.com/gin-gonic/gin"
	"web/internal/comment"
	"web/internal/user"
)

func RegisterRoutes(r *gin.Engine, userRepo user.Repository) {
	{
		v1 := r.Group("/v1")

		userHandler := user.NewHandler(userRepo)

		v1.POST("/users", userHandler.CreateHandler)
		v1.GET("/users", userHandler.GetAllHandler)
		v1.GET("/users/:id", userHandler.GetByIDHandler)
		v1.PUT("/users/:id", userHandler.UpdateAllHandler)
		v1.PATCH("/users/:id", userHandler.UpdatePartialHandler)
		v1.DELETE("/users/:id", userHandler.DeleteHandler)

		v1.POST("/comments", comment.CreateHandler)
	}
}
