package api

import (
	"github.com/gin-gonic/gin"
	"web/internal/comment"
	"web/internal/user"
)

func RegisterRoutes(r *gin.Engine) {
	{
		v1 := r.Group("/v1")

		v1.POST("/users", user.CreateHandler)
		v1.GET("/users", user.GetAllHandler)
		v1.GET("/users/:id", user.GetByIDHandler)
		v1.PUT("/users/:id", user.UpdateAllHandler)
		v1.PATCH("/users/:id", user.UpdatePartialHandler)
		v1.DELETE("/users/:id", user.DeleteHandler)

		v1.POST("/comments", comment.CreateHandler)
	}
}
