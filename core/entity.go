package core

import (
	"time"
)

type User struct {
	ID             int       `json:"id"`
	UnixID         string    `json:"unix_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Country        string    `json:"country"`
	Addreas        string    `json:"addreas"`
	BioUser        string    `json:"bio_user"`
	FBLink         string    `json:"fb_link"`
	IGLink         string    `json:"ig_link"`
	PasswordHash   string    `json:"password_hash"`
	StatusAccount  string    `json:"status_account"`
	AvatarFileName string    `json:"avatar_file_name"`
	Token          string    `json:"token"`
	UpdateIDAdmin  string    `json:"update_id_admin"`
	UpdateAtAdmin  time.Time `json:"update_at_admin"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type NotifInvestor struct {
	ID             int       `json:"id"`
	UserInvestorId string    `json:"user_investor_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Document       string    `json:"document"`
	TypeError      string    `json:"type_error"`
	StatusNotif    int       `json:"status_notif"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
