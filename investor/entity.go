package investor

import (
	"time"
)

type User struct {
	ID             int
	UnixID         string
	Name           string
	Email          string
	Phone          int
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
