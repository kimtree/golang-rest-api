package user

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*User) error
	GetAll(offset int, limit int) ([]User, error)
	GetByID(uint) (*User, error)
	GetByEmail(string) (*User, error)
	Update(*User) (*User, error)
	Delete(id uint) error
}

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *DBRepository) GetAll(offset, limit int) ([]User, error) {
	var users []User
	err := r.db.Preload("Comments").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (r *DBRepository) GetByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *DBRepository) GetByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *DBRepository) Update(u *User) (*User, error) {
	return u, r.db.Save(u).Error
}

func (r *DBRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
