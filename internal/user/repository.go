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

func GetByID(id uint) (User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	return user, err
}

func Update(u *User) (*User, error) {
	return u, db.DB.Save(u).Error
}

func Delete(id uint) error {
	return db.DB.Delete(&User{}, id).Error
}
