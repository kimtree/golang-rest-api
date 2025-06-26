package user

import (
	"web/internal/comment"
)

type User struct {
	ID       uint              `gorm:"primaryKey" json:"id"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Comments []comment.Comment `json:"comments,omitempty"`
}
