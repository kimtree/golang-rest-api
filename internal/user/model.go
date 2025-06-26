package user

import (
	"web/internal/comment"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint              `gorm:"primaryKey" json:"id"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Comments []comment.Comment `json:"comments,omitempty"`
}
