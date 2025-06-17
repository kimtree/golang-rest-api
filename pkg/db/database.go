package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(endpoint string) {
	database, err := gorm.Open(sqlite.Open(endpoint), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	DB = database
}

func Migrate(model interface{}) error {
	return DB.AutoMigrate(model)
}
