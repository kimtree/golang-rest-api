package comment

import (
	"web/internal/pkg/db"
)

func Create(c *Comment) error {
	return db.DB.Create(c).Error
}
