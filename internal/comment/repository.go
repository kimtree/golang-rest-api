package comment

import "web/pkg/db"

func Create(c *Comment) error {
	return db.DB.Create(c).Error
}
