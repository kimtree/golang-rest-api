package user

import "web/pkg/db"

func Create(u *User) error {
	return db.DB.Create(u).Error
}

func GetAll() ([]User, error) {
	var users []User
	err := db.DB.Find(&users).Error
	return users, err
}

func GetByID(id string) (User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	return user, err
}
