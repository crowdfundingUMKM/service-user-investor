package core

import (
	"time"
)

type User struct {
	ID             int
	UnixID         string
	Name           string
	Email          string
	Phone          string
	Country        string
	Addreas        string
	BioUser        string
	FBLink         string
	IGLink         string
	PasswordHash   string
	StatusAccount  string
	AvatarFileName string
	Token          string
	UpdateIDAdmin  string
	UpdateAtAdmin  time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
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
