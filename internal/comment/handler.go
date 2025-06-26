package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a copmment
// @Description This API will create a comment
// @Tags        comments
// @Accept json
// @Produce json
// @Param       request body CreateCommentReq true "Comment"
// @Success 201 {object} Comment
// @Router /comments [post]
func CreateHandler(c *gin.Context) {
	var req CreateCommentReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := Comment{
		UserID: req.UserID,
		Text:   req.Text,
	}

	if err := Create(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}
