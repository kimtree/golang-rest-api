package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userRepo Repository) {
	{
		v1 := r.Group("/v1")

		userHandler := NewHandler(userRepo)

		v1.POST("/users", userHandler.CreateHandler)
		v1.GET("/users", userHandler.GetAllHandler)
		v1.GET("/users/:id", userHandler.GetByIDHandler)
		v1.PUT("/users/:id", userHandler.UpdateAllHandler)
		v1.PATCH("/users/:id", userHandler.UpdatePartialHandler)
		v1.DELETE("/users/:id", userHandler.DeleteHandler)
	}
}
