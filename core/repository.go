package core

import (
	"gorm.io/gorm"
)

// KONTRAK

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByPhone(phone string) (User, error)
	FindByUnixID(unix_id string) (User, error)
	Update(user User) (User, error)
	UploadAvatarImage(user User) (User, error)
	UpdatePassword(user User) (User, error)
	UpdateStatusAccount(user User) (User, error)
	UpdateToken(user User) (User, error)
	GetAllUser() ([]User, error)
	DeleteUser(user User) (User, error)

	SaveReport(report NotifInvestor) (NotifInvestor, error)
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

func (r *repository) FindByUnixID(unix_id string) (User, error) {
	var user User

	err := r.db.Where("unix_id = ?", unix_id).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Model(&user).Updates(User{Name: user.Name, Phone: user.Phone, BioUser: user.BioUser, Addreas: user.Addreas, Country: user.Country, FBLink: user.FBLink, IGLink: user.IGLink}).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UploadAvatarImage(user User) (User, error) {
	err := r.db.Model(&user).Updates(User{AvatarFileName: user.AvatarFileName}).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// make update just row status_account admin request
func (r *repository) UpdateStatusAccount(user User) (User, error) {
	// update status_account and update_by_admin

	err := r.db.Model(&user).Updates(User{
		StatusAccount: user.StatusAccount,
		UpdateIDAdmin: user.UpdateIDAdmin,
	}).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// UpdateToken
func (r *repository) UpdateToken(user User) (User, error) {
	err := r.db.Model(&user).Update("token", user.Token).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// update password
func (r *repository) UpdatePassword(user User) (User, error) {
	err := r.db.Model(&user).Update("password_hash", user.PasswordHash).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// get all user
func (r *repository) GetAllUser() ([]User, error) {
	var user []User

	err := r.db.Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// delete user
func (r *repository) DeleteUser(user User) (User, error) {
	err := r.db.Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// REPORT

func (r *repository) SaveReport(notifInvestor NotifInvestor) (NotifInvestor, error) {
	err := r.db.Create(&notifInvestor).Error
	if err != nil {
		return notifInvestor, err
	}
	return notifInvestor, nil

}
