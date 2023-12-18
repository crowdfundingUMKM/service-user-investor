package core

import (
	"time"
)

// type User struct {
// 	ID             int
// 	UnixID         string
// 	Name           string
// 	Email          string
// 	Phone          string
// 	PasswordHash   string
// 	BioUser        string
// 	AvatarFileName string
// 	StatusAccount  string
// 	Token          string
// 	UpdateByAdmin  string
// 	CreatedAt      time.Time
// 	UpdatedAt      time.Time
// }

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
	ID             int
	UserInvestorId string
	Title          string
	Description    string
	TypeInfo       string
	Document       string
	StatusNotif    int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
