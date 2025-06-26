package comment

type Comment struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Text   string `json:"text"`
	UserID uint
}
