package user

import (
	"errors"
	"gorm.io/gorm"
	"web/pkg/db"
)

func Create(u *User) error {
	return db.DB.Create(u).Error
}

func GetAll(offset, limit int) ([]User, error) {
	var users []User
	err := db.DB.Preload("Comments").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func GetByID(id uint) (*User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	return &user, err
}

func GetByEmail(email string) (*User, error) {
	var user User
	err := db.DB.Where("email = ?", email).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func Update(u *User) (*User, error) {
	return u, db.DB.Save(u).Error
}

func Delete(id uint) error {
	return db.DB.Delete(&User{}, id).Error
}
