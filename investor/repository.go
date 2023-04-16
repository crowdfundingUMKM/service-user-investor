package investor

import (
	"gorm.io/gorm"
)

// KONTRAK

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByPhone(phone string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) FindByPhone(phone string) (User, error) {
	var user User

	err := r.db.Where("phone = ?", phone).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil

}
