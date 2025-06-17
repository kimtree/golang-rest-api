package api

import (
	"github.com/gin-gonic/gin"
	"web/internal/user"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/users", user.CreateHandler)
	r.GET("/users", user.GetAllHandler)
	r.GET("/users/:id", user.GetByIDHandler)
}
