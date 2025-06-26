package comment

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	{
		v1 := r.Group("/v1")

		v1.POST("/comments", CreateHandler)
	}
}
